package computing

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/swanchain/go-computing-provider/internal/models"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

const resourceConfigFile = "price.toml"

var resourcePrice = `TARGET_CPU="0.2"          # SWAN/thread-hour
TARGET_MEMORY="0.1"       # SWAN/GB-hour
TARGET_HD_EPHEMERAL="0.005" # SWAN/GB-hour
TARGET_GPU_DEFAULT="1.6"  # SWAN/Default GPU unit a hour
`

var resourcePriceDefault = map[string]string{
	"TARGET_CPU":          "0.2",
	"TARGET_MEMORY":       "0.1",
	"TARGET_HD_EPHEMERAL": "0.005",
	"TARGET_GPU_DEFAULT":  "1.6",
}

func GeneratePriceConfig() error {
	cpRepoPath, _ := os.LookupEnv("CP_PATH")
	resourcePriceFile := filepath.Join(cpRepoPath, resourceConfigFile)
	if _, err := os.Stat(resourcePriceFile); err == nil {
		return fmt.Errorf("price.conf configuration already exists")
	}

	file, err := os.OpenFile(resourcePriceFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("failed to open file: %s \n", resourcePriceFile)
	}
	defer file.Close()

	file.WriteString(resourcePrice)

	var statisticalSources []*models.NodeResource
	k8sService := NewK8sService()
	if k8sService != nil && k8sService.k8sClient != nil {
		statisticalSources, err = k8sService.StatisticalSources(context.TODO())
		if err != nil {
			return fmt.Errorf("failed to get gpu resource from k8s, error: %v", err)
		}
	} else {
		dockerService := NewDockerService()
		containerLogStr, err := dockerService.ContainerLogs("resource-exporter")
		if err != nil {
			return fmt.Errorf("failed to get gpu resource from docker, error: %v", err)
		}
		var nodeResource models.NodeResource
		if err = json.Unmarshal([]byte(containerLogStr), &nodeResource); err != nil {
			return err
		}
		statisticalSources = append(statisticalSources, &nodeResource)
	}

	var gpuMap = make(map[string]string)
	for _, source := range statisticalSources {
		for _, detail := range source.Gpu.Details {
			gpuMap[detail.ProductName] = detail.ProductName
		}
	}

	for _, gpu := range gpuMap {
		gpuStr := strings.ReplaceAll(gpu, "NVIDIA ", "")
		data := fmt.Sprintf("TARGET_GPU_%s=\"\" # SWAN/%s GPU unit a hour", strings.ReplaceAll(gpuStr, " ", "_"), strings.ReplaceAll(gpuStr, " ", "_"))
		file.WriteString(data)
	}
	fmt.Printf("Successfully generated resource price configuration file at %s \n", resourcePriceFile)
	return nil
}

func ReadPriceConfig() (HardwarePrice, error) {
	var hardwarePrice HardwarePrice
	hardwarePrice.GpusPrice = make(map[string]string)
	cpRepoPath, _ := os.LookupEnv("CP_PATH")

	var priceConfig map[string]string
	_, err := os.Stat(filepath.Join(cpRepoPath, resourceConfigFile))
	if err != nil {
		logs.GetLogger().Warnf("no price configured, use default price")
		priceConfig = resourcePriceDefault
	} else {
		var rc Config
		_, err = toml.DecodeFile(filepath.Join(cpRepoPath, resourceConfigFile), &rc.Resources)
		if err != nil {
			return hardwarePrice, err
		}
		priceConfig = rc.Resources
	}

	for key, value := range priceConfig {
		switch key {
		case "TARGET_CPU":
			hardwarePrice.TARGET_CPU = value
		case "TARGET_MEMORY":
			hardwarePrice.TARGET_MEMORY = value
		case "TARGET_HD_EPHEMERAL":
			hardwarePrice.TARGET_HD_EPHEMERAL = value
		case "TARGET_GPU_DEFAULT":
			hardwarePrice.TARGET_GPU_DEFAULT = value
		default:
			hardwarePrice.GpusPrice[key] = value
		}
	}

	return hardwarePrice, nil
}

func GetStructByTag(v interface{}) ([]HardwareField, error) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input param is not a struct")
	}

	var fields []HardwareField
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("tag")

		if tag != "" {
			tagValue, err := strconv.Atoi(tag)
			if err != nil {
				continue
			}
			fields = append(fields, HardwareField{
				TagValue: tagValue,
				Name:     field.Name,
				Value:    val.Field(i).String(),
			})
		}
	}

	sort.Slice(fields, func(i, j int) bool {
		return fields[i].TagValue < fields[j].TagValue
	})

	return fields, nil
}

type HardwareField struct {
	TagValue int
	Name     string
	Value    string
}

type Config struct {
	Resources map[string]string `toml:"-"`
}

type HardwarePrice struct {
	TARGET_CPU          string `toml:"TARGET_CPU" tag:"1"`
	TARGET_MEMORY       string `toml:"TARGET_MEMORY" tag:"2"`
	TARGET_HD_EPHEMERAL string `toml:"TARGET_HD_EPHEMERAL" tag:"3"`
	TARGET_GPU_DEFAULT  string `toml:"TARGET_GPU_DEFAULT" tag:"4"`
	GpusPrice           map[string]string
}
