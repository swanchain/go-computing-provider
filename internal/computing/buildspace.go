package computing

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/util"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	yamlDeployName = "deploy.yaml"
	ymlDeployName  = "deploy.yml"
	modelSetName   = "model-setting.json"
)

func DownloadSpaceResources(jobUuid string, files []models.SpaceFile) (DeployParam, error) {
	updateJobStatus(jobUuid, models.DEPLOY_DOWNLOAD_SOURCE)
	var deployParam DeployParam

	var err error
	cpRepoPath, _ := os.LookupEnv("CP_PATH")
	buildFolder := filepath.Join(cpRepoPath, "build")
	if len(files) > 0 {
		var containsYaml bool
		var yamlName string
		var modelsSettingFileName string

		var fileNames []string
		for _, file := range files {
			fileNames = append(fileNames, file.Name)
			dirPath := filepath.Dir(file.Name)
			if err = os.MkdirAll(filepath.Join(buildFolder, dirPath), os.ModePerm); err != nil {
				return deployParam, err
			}
			if err = downloadFile(filepath.Join(buildFolder, file.Name), file.URL, file.Iv, file.SymmetricKey); err != nil {
				return deployParam, fmt.Errorf("failed to downloading file: %w", err)
			}

			if strings.HasSuffix(strings.ToLower(file.Name), yamlDeployName) ||
				strings.HasSuffix(strings.ToLower(file.Name), ymlDeployName) {
				containsYaml = true
				yamlName = file.Name
			}
			if strings.EqualFold(file.Name, modelSetName) {
				modelsSettingFileName = file.Name
			}
		}
		prefix := commonPrefix(fileNames)
		imagePath := filepath.Join(buildFolder, prefix)

		var modelsSettingFilePath string
		var yamlPath string
		if modelsSettingFileName != "" {
			modelsSettingFilePath = filepath.Join(buildFolder, modelsSettingFileName)
		}
		if yamlName != "" {
			yamlPath = filepath.Join(buildFolder, yamlName)
		}

		deployParam.ContainsYaml = containsYaml
		deployParam.YamlFilePath = yamlPath
		deployParam.BuildImagePath = imagePath
		deployParam.ModelsSettingFilePath = modelsSettingFilePath

		return deployParam, nil
	} else {
		return deployParam, fmt.Errorf("not found the space")
	}
}

func commonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		str := strs[0]
		return str[:strings.LastIndex(str, "/")]
	} else {
		prefix := strs[0]
		for _, str := range strs {
			for !strings.HasPrefix(str, prefix) {
				if len(prefix) == 0 {
					return ""
				}
				prefix = prefix[:len(prefix)-1]
			}
		}
		return prefix
	}
}

func BuildImagesByDockerfile(jobUuid, spaceName, imagePath string) (string, string) {
	updateJobStatus(jobUuid, models.DEPLOY_BUILD_IMAGE)
	spaceFlag := spaceName + jobUuid[strings.LastIndex(jobUuid, "-"):]
	imageName := fmt.Sprintf("lagrange/%s:%d", spaceFlag, time.Now().Unix())
	if conf.GetConfig().Registry.ServerAddress != "" {
		imageName = fmt.Sprintf("%s/%s:%d",
			strings.TrimSpace(conf.GetConfig().Registry.ServerAddress), spaceFlag, time.Now().Unix())
	}
	imageName = strings.ToLower(imageName)

	dockerfilePath := filepath.Join(imagePath, "Dockerfile")
	if _, err := os.Stat(dockerfilePath); err != nil {
		dockerfilePath = filepath.Join(imagePath, "dockerfile")
		if _, err := os.Stat(dockerfilePath); err != nil {
			logs.GetLogger().Errorf("not found Dockerfile, path: %s", dockerfilePath)
			return "", ""
		}
	}

	log.Printf("Image path: %s", imagePath)

	dockerService := NewDockerService()
	if err := dockerService.BuildImage(imagePath, imageName); err != nil {
		logs.GetLogger().Errorf("Error building Docker image: %v", err)
		return "", ""
	}

	if conf.GetConfig().Registry.ServerAddress != "" {
		updateJobStatus(jobUuid, models.DEPLOY_PUSH_IMAGE)
		if err := dockerService.PushImage(imageName); err != nil {
			logs.GetLogger().Errorf("Error Docker push image: %v", err)
			return "", ""
		}
	}
	return imageName, dockerfilePath
}

func downloadFile(filepath, url, iv, symmetricKey string) error {
	os.Remove(filepath)
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("url: %s, unexpected status code: %d", url, resp.StatusCode)
	}

	var reader io.Reader
	if iv != "" && symmetricKey != "" {
		data, err := decryptData(resp.Body, iv, symmetricKey)
		if err != nil {
			return err
		}
		reader = strings.NewReader(string(data))
	} else {
		reader = resp.Body
	}
	_, err = io.Copy(out, reader)
	if err != nil {
		return err
	}

	return nil
}

func decryptData(body io.ReadCloser, ivData, symmetricKeyData string) ([]byte, error) {
	encodedData, err := io.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response, error: %v", err)
	}

	iv, err := base64.StdEncoding.DecodeString(ivData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode iv using base64, error: %v", err)
	}
	symmetricKey, err := base64.StdEncoding.DecodeString(symmetricKeyData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode symmetricKey using base64, error: %v", err)
	}

	cpRepoPath, _ := os.LookupEnv("CP_PATH")
	privateKeyPath := filepath.Join(cpRepoPath, util.RSA_DIR_NAME, util.RSA_PRIVATE_KEY)
	plainSymmetricKey, err := util.DecryptData(privateKeyPath, symmetricKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt symmetricKey using rsa, error: %v", err)
	}

	plainData, err := util.DecryptAES256CFB(encodedData, plainSymmetricKey, iv)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt symmetricKey using aes-256, error: %v", err)
	}
	return plainData, nil
}
