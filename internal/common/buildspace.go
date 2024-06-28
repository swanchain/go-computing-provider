package common

import (
	"errors"
	"fmt"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/models"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var NotFoundError = errors.New("not found resource")

func BuildSpaceTaskImage(spaceUuid string, files []models.SpaceFile) (bool, string, string, string, string, error) {
	var err error
	cpRepoPath, _ := os.LookupEnv("CP_PATH")
	buildFolder := filepath.Join(cpRepoPath, "build")
	if len(files) > 0 {
		for _, file := range files {
			dirPath := filepath.Dir(file.Name)
			if err = os.MkdirAll(filepath.Join(buildFolder, dirPath), os.ModePerm); err != nil {
				return false, "", "", "", "", err
			}
			if err = downloadFile(filepath.Join(buildFolder, file.Name), file.URL); err != nil {
				return false, "", "", "", "", fmt.Errorf("error downloading file: %w", err)
			}
		}

		imagePath := filepath.Join(buildFolder, getDownloadPath(files[0].Name))
		var containsYaml bool
		var yamlPath string
		var modelsSetting string

		err = filepath.WalkDir(imagePath, func(path string, d fs.DirEntry, err error) error {
			if strings.HasSuffix(strings.ToLower(d.Name()), "deploy.yaml") || strings.HasSuffix(strings.ToLower(d.Name()), "deploy.yml") {
				containsYaml = true
				yamlPath = path
			}
			if strings.EqualFold(d.Name(), "model-setting.json") {
				modelsSetting = path
			}
			return nil
		})
		if err != nil {
			return containsYaml, yamlPath, imagePath, modelsSetting, "", err
		}
		return containsYaml, yamlPath, imagePath, modelsSetting, "", nil
	} else {
		logs.GetLogger().Warnf("Space %s is not found.", spaceUuid)
	}
	return false, "", "", "", "", NotFoundError
}

func getDownloadPath(fileName string) string {
	splits := strings.Split(fileName, "/")
	return filepath.Join(splits[0], splits[1], splits[2])
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
