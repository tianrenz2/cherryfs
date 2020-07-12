package context

import (
	"cherryfs/pkg/roles/dir"
	"cherryfs/pkg/comm/pb"
	"cherryfs/pkg/etcd"
	"os"
	"io/ioutil"
	"fmt"
	"log"
	"encoding/json"
	"github.com/google/uuid"
)

type ChunkInfo struct {
	HostId string
	Addr string
	Dirs []dir.Dir
}

type ChunkContext struct {
	HostId string
	MetaAddrs []string
	Address string
	Client  pb.ChunkServerClient
	EtcdCli etcd.EtcdClient
	LcDirs []*dir.Dir
	ResponseId int64
}

type ChunkConfig struct {
	Addr string
	Dirs []string
}


var GlobalChunkCtx *ChunkContext


func (chunkCtx *ChunkContext) ObtainHostId() (error) {
	exist := true
	for exist {
		hostId := uuid.New().String()
		if _, err := chunkCtx.EtcdCli.Get(HostKeyPrefix + "/" + hostId); err != nil {
			chunkCtx.HostId = hostId
			hostIdPath := os.Getenv("HOST_ID_PATH")
			log.Printf("create host id %s\n", hostIdPath)
			f, err := os.Create(hostIdPath)
			_, err = f.WriteString(hostId)

			if err != nil {
				log.Fatalf("failed to save the host id, %v\n", err)
			}

			return nil
		}
	}
	return nil
}

func (chunkCtx *ChunkContext) SetupConfig() (ChunkConfig, error) {
	configPath := os.Getenv("CONF_PATH")
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return ChunkConfig{}, fmt.Errorf("failed to load the config file: %v", err)
	}
	var config = ChunkConfig{}

	err = json.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatalf("failed to read the configration file: %s\n", configPath)
	}

	return config, nil
}

func (chunkCtx *ChunkContext) RegisterChunkService(config ChunkConfig) error {

	lcDirs := make([]dir.Dir, 0)

	for _, d := range chunkCtx.LcDirs {
		lcDirs = append(lcDirs, *d)
	}

	chunkInfo := ChunkInfo{
		HostId:chunkCtx.HostId,
		Addr: chunkCtx.Address,
		Dirs:lcDirs,
	}

	log.Printf("%v\n", chunkInfo)

	infoByte, err := json.Marshal(chunkInfo)

	if err != nil {
		return err
	}

	chunkRegistryKey := fmt.Sprintf("%s/%s", HostKeyPrefix, chunkCtx.HostId)
	log.Printf("putting registry key %s\n", chunkRegistryKey)
	chunkCtx.EtcdCli.Put(chunkRegistryKey, string(infoByte))

	return nil
}
