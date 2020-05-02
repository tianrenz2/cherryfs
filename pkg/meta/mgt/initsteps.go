package main

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
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
	SubgroupNum int
	ConfigHosts []ConfigHost
}

func loadConfig() []ConfigHost {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to load the config file: %v", err)
	}

	var configHosts = make([]ConfigHost, 0)

	json.Unmarshal([]byte(data), &configHosts)

	for _, host := range configHosts {
		log.Println(fmt.Sprintf("Found host: %v", host))
	}
	return configHosts
}



func main() {
	loadConfig()
}