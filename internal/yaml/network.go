package yaml

import (
	"fmt"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"gopkg.in/yaml.v2"
	"net"
	"os"
	"path/filepath"
	"strings"
)

type NetplanConfig struct {
	Network struct {
		Ethernets map[string]EthernetConfig `yaml:"ethernets"`
		Version   int                       `yaml:"version"`
	} `yaml:"network"`
}

type EthernetConfig struct {
	DHCP4       bool     `yaml:"dhcp4,omitempty"`
	Addresses   []string `yaml:"addresses,omitempty"`
	Gateway4    string   `yaml:"gateway4,omitempty"`
	Nameservers struct {
		Addresses []string `yaml:"addresses,omitempty"`
		Search    []string `yaml:"search,omitempty"`
	} `yaml:"nameservers,omitempty"`
}

func GetNetSegment() ([]string, error) {
	netplanDir := "/etc/netplan"
	var config NetplanConfig
	err := filepath.Walk(netplanDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && (filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml") {
			return parseNetplanFile(path, &config)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	var ipSegment []string
	for _, ethernet := range config.Network.Ethernets {
		for _, address := range ethernet.Addresses {
			_, ipNet, err := net.ParseCIDR(address)
			if err != nil {
				logs.GetLogger().Errorf("failed to parse CIDR, address: %s, error: %v", address, err)
				continue
			}
			ipSegment = append(ipSegment, ipNet.String())
		}
	}

	if len(ipSegment) == 0 {
		netAddress, err := getNetAddress()
		if err != nil {
			return nil, err
		}
		ipSegment = netAddress
	}
	return ipSegment, nil
}

func parseNetplanFile(filepath string, config *NetplanConfig) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to read file, name: %s, error: %v", err)
	}
	return yaml.Unmarshal(data, config)
}

func getNetAddress() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var addrList []string
	for _, iface := range interfaces {
		if strings.HasPrefix(iface.Name, "cali") || strings.HasPrefix(iface.Name, "docker") ||
			strings.HasPrefix(iface.Name, "kube") || strings.HasPrefix(iface.Name, "lo") ||
			strings.Contains(iface.Name, "calico") {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			logs.GetLogger().Errorf("failed to get net interface, interfaceName: %s", iface.Name)
			continue
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}
			if ipNet.IP.To4() != nil {
				ipStr := addr.String()
				addrList = append(addrList, ipNet.IP.Mask(ipNet.Mask).String()+ipStr[strings.Index(ipStr, "/"):])
			}
		}
	}
	return addrList, nil
}
