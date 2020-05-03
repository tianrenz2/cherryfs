package mgt

import (
	"io/ioutil"
	"log"
	"encoding/json"
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

func LoadHosts(configHosts []ConfigHost) []Host {
	var hosts = make([]Host, 0)
	for _, configHost := range configHosts {
		dirs := make([]Dir, 0)
		for _, dir := range configHost.Dirs {
			dirs = append(dirs, Dir{Path:dir})
		}

		hosts = append(hosts, Host{
			Hostname: configHost.Hostname,
			Address: configHost.Address,
			Dirs:dirs,
		})
	}
	return hosts
}
