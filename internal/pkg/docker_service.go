package pkg

import (
	"archive/tar"
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types/filters"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/swanchain/go-computing-provider/conf"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

const BuildFileName = "build.log"

type DockerService struct {
	c *client.Client
}

func NewDockerService() *DockerService {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err.Error())
	}
	return &DockerService{
		c: cli,
	}
}

func ExtractExposedPort(dockerfilePath string) (string, error) {
	file, err := os.Open(dockerfilePath)
	if err != nil {
		return "", fmt.Errorf("unable to open Dockerfile: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	exposedPort := ""
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "EXPOSE") {
			re := regexp.MustCompile(`\d+`)
			exposedPort = re.FindString(line)
			break
		}
	}

	if exposedPort == "" {
		return "", fmt.Errorf("no exposed port found in Dockerfile")
	}

	return exposedPort, nil
}
func RunContainer(imageName, dockerfilePath string) string {
	exposedPort, err := ExtractExposedPort(dockerfilePath)
	if err != nil {
		log.Printf("Failed to extract exposed port: %v", err)
		return ""
	}

	portMapping := exposedPort + ":" + exposedPort
	err = RemoveContainerIfExists(imageName)
	if err != nil {
		log.Printf("Failed to remove existing container: %v", err)
		return ""
	}
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("docker", "run", "-d", "-p", portMapping, imageName)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		log.Printf("run container error: %v\n%s", err, stderr.String())
		return ""
	}

	containerID := strings.TrimSpace(stdout.String())

	// Clear the stdout buffer
	stdout.Reset()

	cmd = exec.Command("docker", "port", containerID)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		log.Printf("get container port error: %v\n%s", err, stderr.String())
		return ""
	}

	portMapping = strings.TrimSpace(stdout.String())
	fmt.Printf("Port mapping: %s\n", portMapping)

	re := regexp.MustCompile(`0\.0\.0\.0:(\d+)`)
	match := re.FindStringSubmatch(portMapping)
	if len(match) < 2 {
		log.Printf("unexpected port mapping format: %s", portMapping)
		return ""
	}

	hostPort := match[1]

	// Replace "0.0.0.0" with the desired IP address (e.g., "127.0.0.1")
	hostIP := "127.0.0.1"

	url := "http://" + hostIP + ":" + hostPort
	return url
}

func RemoveContainerIfExists(imageName string) error {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("docker", "ps", "-a", "-q", "--filter", "ancestor="+imageName)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("list containers error: %v\n%s", err, stderr.String())
		return err
	}

	containerIDs := strings.Split(strings.TrimSpace(stdout.String()), "\n")
	if len(containerIDs) == 0 || containerIDs[0] == "" {
		log.Printf("No container with image %s found.", imageName)
		return nil
	}

	for _, containerID := range containerIDs {
		stdout.Reset()
		stderr.Reset()

		cmd = exec.Command("docker", "rm", "-f", containerID)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err = cmd.Run()
		if err != nil {
			log.Printf("remove container error: %v\n%s", err, stderr.String())
			return err
		}

		log.Printf("Removed container with ID %s", containerID)
	}

	return nil
}

func (ds *DockerService) BuildImage(buildPath, imageName string) error {
	// Create a buffer
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	filepath.Walk(buildPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(buildPath, path)
		if err != nil {
			return err
		}
		header.Name = relPath
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			if _, err := io.Copy(tw, file); err != nil {
				return err
			}
		}
		return nil
	})

	dockerFileTarReader := bytes.NewReader(buf.Bytes())
	buildResponse, err := ds.c.ImageBuild(context.Background(), dockerFileTarReader, types.ImageBuildOptions{
		Context: dockerFileTarReader,
		Tags:    []string{imageName},
	})
	if err != nil {
		return err
	}
	defer buildResponse.Body.Close()

	logFile, err := os.Create(filepath.Join(buildPath, BuildFileName))
	if err != nil {
		return err
	}
	defer logFile.Close()

	logWriters := []io.Writer{logFile, os.Stdout}
	multiWriter := io.MultiWriter(logWriters...)

	_, err = io.Copy(multiWriter, buildResponse.Body)
	if err != nil {
		return err
	}
	return nil
}

type ErrorLine struct {
	Error       string `json:"error"`
	ErrorDetail struct {
		Message string `json:"message"`
	} `json:"errorDetail"`
}

func (ds *DockerService) PushImage(imagesName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*6000)
	defer cancel()

	var authConfig = types.AuthConfig{
		ServerAddress: conf.GetConfig().Registry.ServerAddress,
		Username:      conf.GetConfig().Registry.UserName,
		Password:      conf.GetConfig().Registry.Password,
	}
	authConfigBytes, _ := json.Marshal(authConfig)
	authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)
	opts := types.ImagePushOptions{RegistryAuth: authConfigEncoded}

	// retry
	retries := 5
	var err error
	for i := 0; i < retries; i++ {
		rd, rerr := ds.c.ImagePush(ctx, imagesName, opts)
		if rerr == nil {
			err = printOut(rd)
			rd.Close()
			if err == nil {
				return nil
			}
		} else {
			err = rerr
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}

func printOut(rd io.Reader) error {
	var lastLine string
	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		lastLine = scanner.Text()
		println(scanner.Text())
	}
	errLine := &ErrorLine{}
	json.Unmarshal([]byte(lastLine), errLine)
	if errLine.Error != "" {
		return errors.New(errLine.Error)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (ds *DockerService) RemoveImage(imageId string) error {
	ctx := context.Background()
	_, err := ds.c.ImageRemove(ctx, imageId, types.ImageRemoveOptions{
		Force:         true,
		PruneChildren: true,
	})
	return err
}

func (ds *DockerService) CleanResource() {
	images, err := ds.c.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		ulog.Errorf("Failed get image list, error: %+v", err)
		return
	}

	for _, image := range images {
		var specialImage bool
		for _, tag := range image.RepoTags {
			if strings.HasPrefix(tag, "filswan/ubi-worker") || strings.HasPrefix(tag, "filswan/cpu-model-collector") {
				specialImage = true
				break
			}
		}

		if !specialImage {
			ds.RemoveImage(image.ID)
		}
	}

	ctx := context.Background()
	danglingFilters := filters.NewArgs()
	danglingFilters.Add("dangling", "true")
	_, err = ds.c.ImagesPrune(ctx, danglingFilters)
	if err != nil {
		ulog.Errorf("Failed delete dangling image, error: %+v", err)
		return
	}

	if _, err = ds.c.ContainersPrune(ctx, filters.NewArgs()); err != nil {
		ulog.Errorf("Failed delete unused container, error: %+v", err)
		return
	}
}
