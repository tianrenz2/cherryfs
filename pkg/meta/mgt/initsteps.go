package mgt

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"cherryfs/pkg/roles/host"
	"cherryfs/pkg/roles/dir"
)

const (
	configPath = "../../../cluster_config.json"
)

type ConfigHost struct {
	Hostname string
	Address string
	Dirs []string
}

type Config struct {
	Subgroupnum int
	Hosts []ConfigHost
}

func LoadConfig() Config {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to load the config file: %v", err)
	}
	var config = Config{}

	e := json.Unmarshal([]byte(data), &config)

	if err != nil {
		log.Fatalf("%v", e)
	}
	return config
}

func LoadHosts(configHosts []ConfigHost) []host.Host {
	var hosts = make([]host.Host, 0)
	for _, configHost := range configHosts {
		dirs := make([]dir.Dir, 0)
		for _, d := range configHost.Dirs {
			dirs = append(dirs, dir.Dir {Path:d})
		}

		hosts = append(hosts, host.Host{
			Hostname: configHost.Hostname,
			Address: configHost.Address,
			Dirs:dirs,
		})
	}
	return hosts
}
