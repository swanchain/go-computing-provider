package computing

import (
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
	"strings"
)

const resourceConfigFile = "price.toml"

var resourcePrice = `
TARGET_CPU=""          # SWAN/thread-hour
TARGET_MEMORY=""       # SWAN/GB-hour
TARGET_HD_EPHEMERAL="" # SWAN/GB-hour
TARGET_HD_PERS_HDD=""  # SWAN/GB-hour
TARGET_HD_PERS_SSD=""  # SWAN/GB-hour
TARGET_HD_PERS_NVME="" # SWAN/GB-hour
TARGET_GPU_DEFAULT=""   # SWAN/Default GPU unit a hour
`

func GeneratePriceConfig() error {
	cpRepoPath, _ := os.LookupEnv("CP_PATH")
	resourcePriceFile := filepath.Join(cpRepoPath, resourceConfigFile)
	if _, err := os.Stat(resourcePriceFile); err == nil {
		return fmt.Errorf("price.conf configuration already exists")
	}

	os.Create(resourcePriceFile)
	os.WriteFile(filepath.Join(cpRepoPath, resourceConfigFile), []byte(resourcePrice), 0666)

	statisticalSources, err := NewK8sService().StatisticalSources(context.TODO())
	if err != nil {
		return fmt.Errorf("failed to get gpu resource, error: %v", err)
	}

	var gpuMap = make(map[string]string)
	for _, source := range statisticalSources {
		for _, detail := range source.Gpu.Details {
			gpuMap[detail.ProductName] = detail.ProductName
		}
	}

	for _, gpu := range gpuMap {
		data := fmt.Sprintf("TARGET_GPU_%s=\"\" # SWAN/%s GPU unit a hour", strings.ReplaceAll(gpu, "NVIDIA ", ""), strings.ReplaceAll(gpu, "NVIDIA ", ""))
		os.WriteFile(filepath.Join(cpRepoPath, resourceConfigFile), []byte(data), 0666)
	}
	fmt.Printf("Successfully generated resource price configuration file at %s", resourcePriceFile)
	return nil
}

func ReadPriceConfig() error {
	cpRepoPath, _ := os.LookupEnv("CP_PATH")
	if _, err := os.Stat(filepath.Join(cpRepoPath, resourceConfigFile)); err != nil {
		return err
	}

	var config Config
	_, err := toml.DecodeFile(filepath.Join(cpRepoPath, resourceConfigFile), &config.Resources)
	if err != nil {
		return err
	}

	for key, value := range config.Resources {
		fmt.Printf("%s = %s\n", key, value)
	}
	return nil
}

type Config struct {
	Resources map[string]string `toml:"-"`
}
