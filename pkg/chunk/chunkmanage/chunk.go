package chunkmanage

import (
	"cherryfs/pkg/context"
	"cherryfs/pkg/role/dir"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func StartupChunk() error {

	metaAddrs := strings.Split(os.Getenv("ETCDADDR"), ",")
	context.GlobalChunkCtx.MetaAddrs = metaAddrs

	hostFile, err := os.Open(os.Getenv("HOSTIDPATH"))

	if os.IsNotExist(err) {
		context.GlobalChunkCtx.ObtainHostId()
		log.Printf("generate a new host id: %s\n", context.GlobalChunkCtx.HostId)
	} else {
		b, _ := ioutil.ReadAll(hostFile)
		log.Printf("existing hostid: %s\n", string(b))
		context.GlobalChunkCtx.HostId = string(b)
	}

	chunkCfg, err := context.GlobalChunkCtx.SetupConfig()
	if err != nil {
		log.Fatalf("failed to setup configuration: %v\n", err)
	}

	for _, d := range chunkCfg.Dirs {
		context.GlobalChunkCtx.LcDirs = append(context.GlobalChunkCtx.LcDirs, &dir.Dir{
			Path:       d,
			HostId:     context.GlobalChunkCtx.HostId,
			UsedSpace:  0,
			TotalSpace: 0,
		})
	}
	CollectDirStats()
	context.GlobalChunkCtx.RegisterChunkService(chunkCfg)

	return nil
}
