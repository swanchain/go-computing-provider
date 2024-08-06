package yaml

import (
	"gopkg.in/errgo.v2/errors"
	corev1 "k8s.io/api/core/v1"
	"strings"
)

const (
	ServiceTypeNodePort = "node-port"
)

type DeployYamlV2 struct {
	Version     string                `yaml:"version"`
	ServiceType string                `yaml:"type"`
	Services    map[string]Service    `yaml:"services"`
	Profiles    Profiles              `yaml:"profiles"`
	Deployment  map[string]Deployment `yaml:"deployment"`
}

func (dy *DeployYamlV2) checkRequired() error {
	if len(dy.Services) <= 0 {
		return errors.New("at least one service must be defined")
	}
	return nil
}

func (dy *DeployYamlV2) ServiceToK8sResource() ([]ContainerResource, error) {
	if err := dy.checkRequired(); err != nil {
		return nil, err
	}
	var containers []ContainerResource
	var waitDelete []string

	for name, deployment := range dy.Deployment {
		containerNew := new(ContainerResource)
		if service, ok := dy.Services[name]; ok {
			containerNew.Name = name
			containerNew.ServiceType = dy.ServiceType

			var depends []ContainerResource
			for _, depend := range service.DependsOn {
				if service, ok := dy.Services[depend]; ok {
					container := new(ContainerResource)
					container.Name = depend
					container.ImageName = service.Image
					if len(service.Command) > 0 {
						container.Command = service.Command
					}
					if len(service.Args) > 0 {
						container.Args = service.Args
					}
					if len(service.Env) > 0 {
						var envVars []corev1.EnvVar
						for _, env := range service.Env {
							envSplit := strings.Split(strings.TrimSpace(env), "=")
							envVars = append(envVars, corev1.EnvVar{
								Name:  envSplit[0],
								Value: envSplit[1],
							})
						}
						container.Env = envVars
					}
					if len(service.Expose) > 0 {
						var ports []corev1.ContainerPort
						for _, expose := range service.Expose {
							ports = append(ports, corev1.ContainerPort{
								ContainerPort: int32(expose.Port),
								Protocol:      getProtocol(expose.Protocol),
							})
						}
						container.Ports = ports
					}

					if service.Config.Name != "" && service.Config.Path != "" {
						container.VolumeMounts = ConfigFile{
							Name: service.Config.Name,
							Path: service.Config.Path,
						}
					}

					if len(service.ReadyCmd) > 0 {
						container.ReadyCmd = service.ReadyCmd
					}

					if deployment.Akash.Count != 0 {
						container.Count = deployment.Akash.Count
					}
					if deployment.Lagrange.Count != 0 {
						container.Count = deployment.Lagrange.Count
					}
					depends = append(depends, *container)
					waitDelete = append(waitDelete, depend)
				}
			}
			containerNew.Depends = depends
			containerNew.ImageName = service.Image
			if len(service.Command) > 0 {
				containerNew.Command = service.Command
			}
			if len(service.Args) > 0 {
				containerNew.Args = service.Args
			}
			if len(service.Env) > 0 {
				var envVars []corev1.EnvVar
				for _, env := range service.Env {
					envSplit := strings.Split(strings.TrimSpace(env), "=")
					envVars = append(envVars, corev1.EnvVar{
						Name:  envSplit[0],
						Value: envSplit[1],
					})
				}
				containerNew.Env = envVars
			}
			if len(service.Expose) > 0 {
				var ports []corev1.ContainerPort
				for _, expose := range service.Expose {
					ports = append(ports, corev1.ContainerPort{
						ContainerPort: int32(expose.Port),
						Protocol:      getProtocol(expose.Protocol),
						HostPort:      int32(expose.As),
					})
				}
				containerNew.Ports = ports
			}

			if service.Config.Name != "" && service.Config.Path != "" {
				containerNew.VolumeMounts = ConfigFile{
					Name: service.Config.Name,
					Path: service.Config.Path,
				}
			}
			containerNew.Models = service.Models
		}

		containerNew.ResourceLimit = make(corev1.ResourceList)
		if deployment.Akash.Count != 0 {
			containerNew.Count = deployment.Akash.Count
		}
		if deployment.Lagrange.Count != 0 {
			containerNew.Count = deployment.Lagrange.Count
		}
		containers = append(containers, *containerNew)
	}

	var result []ContainerResource
	for _, c := range containers {
		var flag bool
		for _, needToDel := range waitDelete {
			if c.Name == needToDel {
				flag = true
				break
			}
		}
		if !flag {
			result = append(result, c)
		}
	}

	return result, nil
}

type Service struct {
	Name      string
	Image     string   `yaml:"image"`
	Command   []string `yaml:"command"`
	Args      []string `yaml:"args"`
	Env       []string `yaml:"env"`
	Expose    []Expose `yaml:"expose"`
	DependsOn []string `yaml:"depends-on"`
	Config    struct {
		Name string `yaml:"name"`
		Path string `yaml:"path"`
	} `yaml:"config"`
	ReadyCmd []string        `yaml:"ready-cmd"`
	Models   []ModelResource `yaml:"models"`
}

type Expose struct {
	Port int `yaml:"port"`
	To   []struct {
		Global bool `yaml:"global"`
	} `yaml:"to"`
	As       int    `yaml:"as"`
	Protocol string `yaml:"protocol"`
}

type Profiles struct {
	Compute map[string]Compute `yaml:"compute"`
}

type Compute struct {
	Resources struct {
		Cpu struct {
			Units string `yaml:"units"`
		} `yaml:"cpu"`
		Memory struct {
			Size string `yaml:"size"`
		} `yaml:"memory"`
		Storage struct {
			Size string `yaml:"size"`
		} `yaml:"storage"`
		Gpu struct {
			Model string `yaml:"model"`
			Units string `yaml:"units"`
			Size  string `yaml:"size"`
		} `yaml:"gpu"`
	} `yaml:"resources"`
}

type Deployment struct {
	Akash struct {
		Profile string `yaml:"profile"`
		Count   int    `yaml:"count"`
	} `yaml:"akash"`
	Lagrange struct {
		Profile string `yaml:"profile"`
		Count   int    `yaml:"count"`
	} `yaml:"lagrange"`
}

func getProtocol(proto string) corev1.Protocol {
	var result corev1.Protocol
	switch proto {
	case "tcp":
		result = corev1.ProtocolTCP
	case "udp":
		result = corev1.ProtocolUDP
	default:
		result = corev1.ProtocolTCP
	}
	return result
}

type ModelResource struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
	Dir  string `yaml:"dir"`
}
