package mgt

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"cherryfs/pkg/roles/host"
)

const (
	configPath = "/root/go/src/cherryfs/cluster_config.json"
)

type Config struct {
	Subgroupnum int
	Hosts []host.ConfigHost
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

func LoadHosts(configHosts []host.ConfigHost) []host.Host {
	var hosts = make([]host.Host, 0)
	for _, configHost := range configHosts {
		dirs := make([]string, 0)
		for _, d := range configHost.Dirs {
			dirs = append(dirs, d)
		}

		hosts = append(hosts, host.Host{
			Hostname: configHost.Hostname,
			Address: configHost.Address,
			Dirs:dirs,
		})
	}
	return hosts
}
