package main

import (
	_ "embed"
	"fmt"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/internal/yaml"
	"github.com/urfave/cli/v2"
	yaml2 "gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"strings"
)

var networkCmd = &cli.Command{
	Name:  "network",
	Usage: "Manage cp network policy",
	Subcommands: []*cli.Command{
		generateNetworkCmd,
		viewCmd,
	},
	Before: func(c *cli.Context) error {
		cpRepoPath, _ := os.LookupEnv("CP_PATH")
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}
		return nil
	},
}

//go:embed network-policy.yaml
var networkPolicyContent string

var generateNetworkCmd = &cli.Command{
	Name:  "generate",
	Usage: "Generate network policy for k8s cluster",
	Action: func(cctx *cli.Context) error {

		netSegment, err := yaml.GetNetSegment()
		if err != nil {
			return fmt.Errorf("failed to get host subnet mask, error: %v", err)
		}

		type networkSet struct {
			Kind       string `yaml:"kind"`
			APIVersion string `yaml:"apiVersion"`
			Metadata   struct {
				Name   string            `yaml:"name"`
				Labels map[string]string `yaml:"labels"`
			} `yaml:"metadata"`
			Spec struct {
				Nets []string `yaml:"nets"`
			} `yaml:"spec"`
		}
		gns := &networkSet{
			APIVersion: "projectcalico.org/v3",
			Kind:       "GlobalNetworkSet",
			Metadata: struct {
				Name   string            `yaml:"name"`
				Labels map[string]string `yaml:"labels"`
			}{
				Name: models.NetworkNetset,
				Labels: map[string]string{
					"net": models.NetworkNetset,
				},
			},
			Spec: struct {
				Nets []string `yaml:"nets"`
			}{
				Nets: netSegment,
			},
		}

		networkSetBytes, err := yaml2.Marshal(&gns)
		if err != nil {
			return err
		}
		cpRepoPath, _ := os.LookupEnv("CP_PATH")
		filePath := filepath.Join(cpRepoPath, "network-policy.yaml")
		if _, err := os.Stat(filePath); err != nil {
			networkPolicyContent = strings.ReplaceAll(networkPolicyContent, "$NETWORK_SET", string(networkSetBytes))
			networkPolicyContent = strings.ReplaceAll(networkPolicyContent, "$G_SUBNET_NAME", models.NetworkGlobalSubnet)
			networkPolicyContent = strings.ReplaceAll(networkPolicyContent, "$G_OUT_NAME", models.NetworkGlobalOutAccess)
			networkPolicyContent = strings.ReplaceAll(networkPolicyContent, "$G_IN_NAME", models.NetworkGlobalInAccess)
			networkPolicyContent = strings.ReplaceAll(networkPolicyContent, "$G_NS_NAME", models.NetworkGlobalNamespace)
			networkPolicyContent = strings.ReplaceAll(networkPolicyContent, "$G_DNS_NAME", models.NetworkGlobalDns)
			networkPolicyContent = strings.ReplaceAll(networkPolicyContent, "$G_POD_NS_NAME", models.NetworkGlobalPodInNamespace)
			if err = os.WriteFile(filePath, []byte(networkPolicyContent), 0644); err != nil {
				return err
			}
			fmt.Printf("Successfully generated network policy configuration file at %s, please execute this command `kubectl apply -f %s` \n", filePath, filePath)
			return nil
		}
		fmt.Printf("%s already exists \n", filePath)
		return nil
	},
}
