package computing

import (
	"archive/tar"
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/containerd/containerd/namespaces"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/swanchain/go-computing-provider/build"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/swanchain/go-computing-provider/conf"

	"github.com/containerd/containerd"
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

	var authConfig = registry.AuthConfig{
		ServerAddress: conf.GetConfig().Registry.ServerAddress,
		Username:      conf.GetConfig().Registry.UserName,
		Password:      conf.GetConfig().Registry.Password,
	}
	authConfigBytes, _ := json.Marshal(authConfig)
	authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)
	opts := image.PushOptions{RegistryAuth: authConfigEncoded}

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
	_, err := ds.c.ImageRemove(ctx, imageId, image.RemoveOptions{
		Force:         true,
		PruneChildren: true,
	})
	return err
}

func (ds *DockerService) RemoveContainerByName(containerName string) error {
	containerList, err := ds.c.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return err
	}
	for _, c := range containerList {
		for _, name := range c.Names {
			if name == "/"+containerName || strings.Contains(name, containerName) {
				if err := ds.c.ContainerRemove(context.Background(), c.ID, container.RemoveOptions{Force: true}); err != nil {
					return err
				}
				return nil
			}
		}
	}
	return nil
}

func (ds *DockerService) CleanResourceForK8s() {
	containers, err := ds.c.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		logs.GetLogger().Errorf("get all container failed, error: %v", err)
		return
	}

	for _, c := range containers {
		if c.State != "running" {
			ds.c.ContainerRemove(context.Background(), c.ID, container.RemoveOptions{Force: true})
		}
	}

	imagesToKeep := []string{
		build.UBITaskImageIntelCpu,
		build.UBITaskImageIntelGpu,
		build.UBITaskImageAmdCpu,
		build.UBITaskImageAmdGpu,
		build.UBIResourceExporterDockerImage,
	}

	keepSet := make(map[string]bool)
	for _, imageName := range imagesToKeep {
		keepSet[imageName] = true
	}

	allImages, err := ds.c.ImageList(context.Background(), image.ListOptions{})
	if err != nil {
		logs.GetLogger().Errorf("Failed get image list, error: %+v", err)
		return
	}
	for _, img := range allImages {
		for _, tag := range img.RepoTags {
			if !keepSet[tag] {
				ds.c.ImageRemove(context.Background(), tag, image.RemoveOptions{
					Force:         false,
					PruneChildren: true,
				})
			}
		}
	}

	ctx := context.Background()
	danglingFilters := filters.NewArgs()
	danglingFilters.Add("dangling", "true")
	ds.c.ImagesPrune(ctx, danglingFilters)
	ds.c.ContainersPrune(ctx, filters.NewArgs())
}

func (ds *DockerService) CleanResourceForDocker() {
	imagesToKeep := []string{
		build.UBITaskImageIntelCpu,
		build.UBITaskImageIntelGpu,
		build.UBITaskImageAmdCpu,
		build.UBITaskImageAmdGpu,
		build.UBIResourceExporterDockerImage,
	}

	keepSet := make(map[string]bool)
	for _, imageName := range imagesToKeep {
		keepSet[imageName] = true
	}

	allImages, err := ds.c.ImageList(context.Background(), image.ListOptions{})
	if err != nil {
		logs.GetLogger().Errorf("Failed get image list, error: %+v", err)
		return
	}
	for _, img := range allImages {
		for _, tag := range img.RepoTags {
			if !keepSet[tag] {
				ds.c.ImageRemove(context.Background(), tag, image.RemoveOptions{
					Force:         false,
					PruneChildren: true,
				})
			}
		}
	}

	cmd := exec.Command("docker", "system", "prune", "-f")
	if err = cmd.Run(); err != nil {
		logs.GetLogger().Errorf("failed to clean resource, error: %+v", err)
		return
	}
}

func (ds *DockerService) PullImage(imageName string) error {
	filterArgs := filters.NewArgs()
	filterArgs.Add("reference", imageName)

	images, err := ds.c.ImageList(context.Background(), image.ListOptions{
		Filters: filterArgs,
	})
	if err != nil {
		logs.GetLogger().Errorf("get %s image failed, error: %+v", imageName, err)
		return err
	}
	if len(images) > 0 {
		return nil
	} else {
		resp, err := ds.c.ImagePull(context.TODO(), imageName, image.PullOptions{})
		if err != nil {
			return err
		}
		defer resp.Close()
		printOut(resp)
		return nil
	}
}

func (ds *DockerService) CheckRunningContainer(containerName string) (bool, error) {
	containers, err := ds.c.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		logs.GetLogger().Errorf("listing containers failed, error: %v", err)
		return false, err
	}
	containerRunning := false
	for _, c := range containers {
		for _, name := range c.Names {
			if name == "/"+containerName {
				if c.State == "running" {
					containerRunning = true
					break
				}
			}
		}
	}
	if containerRunning {
		return true, nil
	}
	return false, nil
}

func (ds *DockerService) ContainerCreateAndStart(config *container.Config, hostConfig *container.HostConfig, containerName string) error {
	ctx := context.Background()
	resp, err := ds.c.ContainerCreate(ctx, config, hostConfig, nil, nil, containerName)
	if err != nil {
		return err
	}
	return ds.c.ContainerStart(ctx, resp.ID, container.StartOptions{})
}

func (ds *DockerService) ContainerLogs(containerName string) (string, error) {
	ctx := context.Background()
	logReader, err := ds.c.ContainerLogs(ctx, containerName, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     false,
		Tail:       "1",
	})
	if err != nil {
		return "", err
	}
	defer logReader.Close()
	all, err := io.ReadAll(logReader)
	result := string(all)
	index := strings.Index(result, "{")
	if index > 0 {
		return result[index:], nil
	} else {
		return result, nil
	}
}

func (ds *DockerService) GetContainerLogStream(containerName string) (io.ReadCloser, error) {
	ctx := context.Background()
	return ds.c.ContainerLogs(ctx, containerName, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
	})
}

func (ds *DockerService) checkImageExists(imageName string) bool {
	filterArgs := filters.NewArgs()
	filterArgs.Add("reference", imageName)

	images, err := ds.c.ImageList(context.Background(), image.ListOptions{
		Filters: filterArgs,
	})
	if err != nil {
		return false
	}
	return len(images) > 0
}

func (ds *DockerService) IsExistContainer(containerName string) bool {
	containers, err := ds.c.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return false
	}

	exists := false
	for _, c := range containers {
		for _, name := range c.Names {
			if name == "/"+containerName {
				exists = true
				break
			}
		}
		if exists {
			break
		}
	}

	if exists {
		return true
	} else {
		return false
	}
}

func (ds *DockerService) SaveDockerImage(imageName string) (string, error) {
	ctx := context.Background()
	out, err := ds.c.ImageSave(ctx, []string{imageName})
	if err != nil {
		return "", err
	}
	defer out.Close()

	cpRepoPath, _ := os.LookupEnv("CP_PATH")
	imageFolder := filepath.Join(cpRepoPath, "images_cache")
	tarFile := filepath.Join(imageFolder, imageName+".tar")

	os.MkdirAll(filepath.Dir(tarFile), os.ModePerm)
	file, err := os.Create(tarFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.ReadFrom(out)
	if err != nil {
		return "", err
	}
	return tarFile, nil
}

func ImportImageToContainerd(tarFile string) error {
	defer os.Remove(tarFile)
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return err
	}
	defer client.Close()

	file, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer file.Close()

	ctx := namespaces.WithNamespace(context.Background(), "k8s.io")
	img, err := client.Import(ctx, file)
	if err != nil {
		return err
	}
	for _, i := range img {
		fmt.Println("containerd imported image:", i.Name)
	}
	return nil
}
