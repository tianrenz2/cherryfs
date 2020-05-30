package context

import (
	"strings"
	"encoding/json"
	"log"
	"cherryfs/pkg/config"
)


func (ctx *Context) Watcher() {
	watchChan := ctx.EtcdCli.WatchKey(true, HostKeyPrefix)

	for res := range watchChan {
		key := string(res.Events[0].Kv.Key)
		val := res.Events[0].Kv.Value
		addrSlices := strings.Split(key, "/")
		addr := addrSlices[len(addrSlices) - 1]
		var chunkInfo ChunkInfo

		err := json.Unmarshal(val, &chunkInfo)

		if err != nil {
			log.Fatalf("err while parsing data: %v\n", err)
		}

		err = ctx.RegisterChunk(addr, chunkInfo)
		if err != nil {
			log.Fatalf("err while registering %v", err)
		}

		for _, d := range ctx.DManager.Dirs {
			log.Printf("dir size: %d, %d\n", d.UsedSpace, d.TotalSpace)
		}
	}
}

func (ctx *Context) RegisterChunk(addr string, info ChunkInfo) error {
	log.Printf("register chunk %s \n", addr)
	hostId, err := ctx.HManager.AddHost(addr, info.Dirs, ctx.DManager)

	if err != nil {
		return err
	}

	if len(ctx.SGManager.SubGroups) < config.MinReplicaNum {
		newSg, err := ctx.SGManager.InitOneSubgroup(ctx.HManager.Hosts, len(ctx.SGManager.SubGroups), len(ctx.HManager.Hosts) - 1, len(ctx.HManager.Hosts))

		if err != nil {
			return err
		}
		ctx.SGManager.SubGroups = append(ctx.SGManager.SubGroups, newSg)

	} else {
		var minHostSg = ctx.SGManager.SubGroups[0]

		for sgId, _ := range ctx.SGManager.SubGroups {
			sg := ctx.SGManager.GetSubGroupById(sgId)
			if len(sg.Hosts) < len(minHostSg.Hosts) {
				minHostSg = sg
			}
		}

		minHostSg.Hosts = append(minHostSg.Hosts, hostId)
	}

	return nil
}