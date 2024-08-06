package computing

import (
	"errors"
	"fmt"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/models"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var NotFoundError = errors.New("not found resource")

const (
	yamlDeployName = "deploy.yaml"
	ymlDeployName  = "deploy.yml"
	modelSetName   = "model-setting.json"
)

func BuildSpaceTaskImage(spaceUuid string, files []models.SpaceFile) (DeployParam, error) {
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
			if err = downloadFile(filepath.Join(buildFolder, file.Name), file.URL); err != nil {
				return deployParam, fmt.Errorf("error downloading file: %w", err)
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
		logs.GetLogger().Warnf("Space %s is not found.", spaceUuid)
	}
	return deployParam, NotFoundError
}

func commonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

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

func BuildImagesByDockerfile(jobUuid, spaceUuid, spaceName, imagePath string) (string, string) {
	updateJobStatus(jobUuid, models.DEPLOY_BUILD_IMAGE)
	spaceFlag := spaceName + spaceUuid[strings.LastIndex(spaceUuid, "-"):]
	imageName := fmt.Sprintf("lagrange/%s:%d", spaceFlag, time.Now().Unix())
	if conf.GetConfig().Registry.ServerAddress != "" {
		imageName = fmt.Sprintf("%s/%s:%d",
			strings.TrimSpace(conf.GetConfig().Registry.ServerAddress), spaceFlag, time.Now().Unix())
	}
	imageName = strings.ToLower(imageName)
	dockerfilePath := filepath.Join(imagePath, "Dockerfile")
	if dockerfilePath == "" {
		dockerfilePath = filepath.Join(imagePath, "dockerfile")
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

func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("url: %s, unexpected status code: %d", url, resp.StatusCode)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
