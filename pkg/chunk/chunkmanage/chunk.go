package chunkmanage

import (
	"cherryfs/pkg/comm/pb"
	"cherryfs/pkg/roles/dir"
	"os"
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
	"strings"
	"github.com/google/uuid"
	"cherryfs/pkg/etcd"
	"cherryfs/pkg/context"
)

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

func (chunkCtx *ChunkContext) StartupChunk() (error) {
	metaAddrs := strings.Split(os.Getenv("ETCDADDR"), ",")
	chunkCtx.MetaAddrs = metaAddrs
	chunkCtx.ObtainHostId()

	chunkCfg, err := chunkCtx.SetupConfig()
	if err != nil {
		log.Fatalf("failed to setup configuration: %v\n", err)
	}

	for _, d := range chunkCfg.Dirs {
		chunkCtx.LcDirs = append(chunkCtx.LcDirs, &dir.Dir{
			Path: d,
			HostId:chunkCtx.HostId,
			UsedSpace: 0,
			TotalSpace: 0,
		})
	}
	chunkCtx.CollectDirStats()
	chunkCtx.RegisterChunkService(chunkCfg)

	return nil
}

func (chunkCtx *ChunkContext) ObtainHostId() (error) {
	exist := true
	for exist {
		hostId := uuid.New().String()
		if _, err := chunkCtx.EtcdCli.Get(context.HostKeyPrefix + "/" + hostId); err != nil {
			chunkCtx.HostId = hostId
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

	chunkInfo := context.ChunkInfo{
		Addr: chunkCtx.Address,
		Dirs:lcDirs,
	}

	infoByte, err := json.Marshal(chunkInfo)

	if err != nil {
		return err
	}

	chunkRegistryKey := context.HostKeyPrefix + "/" + chunkCtx.HostId
	chunkCtx.EtcdCli.Put(chunkRegistryKey, string(infoByte))

	return nil
}
