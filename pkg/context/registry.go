package context

import (
	"encoding/json"
	"log"
	"cherryfs/pkg/config"
)


func (ctx *Context) RegistryWatcher() {
	// Registry Watcher is watching for new (chunk) services for registration
	watchChan := ctx.EtcdCli.WatchKey(true, HostKeyPrefix)

	for res := range watchChan {
		val := res.Events[0].Kv.Value
		var chunkInfo ChunkInfo

		err := json.Unmarshal(val, &chunkInfo)

		if err != nil {
			log.Fatalf("err while parsing data: %v\n", err)
		}

		err = ctx.RegisterChunk(chunkInfo)
		if err != nil {
			log.Fatalf("err while registering %v", err)
		}

		for _, h := range ctx.HManager.Hosts {
			log.Printf("host: %s\n", h.HostId)
		}

		for _, d := range ctx.DManager.Dirs {
			log.Printf("dir %s size: %d, %d\n", d.Path, d.UsedSpace, d.TotalSpace)
		}
	}
}

func (ctx *Context) RegisterChunk(info ChunkInfo) error {
	log.Printf("register chunk id: %s, addr: %s\n", info.HostId, info.Addr)

	if _, err := ctx.HManager.GetHostByHostId(info.HostId); err == nil {
		log.Printf("chunk %s comes back\n", info.HostId)
		return nil
	}

	hostId, err := ctx.HManager.AddHost(info.HostId, info.Addr, info.Dirs, ctx.DManager)

	if err != nil {
		return err
	}

	// When a new chunk joins in, first check if subgroup number has satisfied the minimum replica number,
	// if not, expand the number of subgroups at first
	if len(ctx.SGManager.SubGroups) < config.MinReplicaNum {
		newSg, err := ctx.SGManager.InitOneSubgroup(ctx.HManager.Hosts, len(ctx.SGManager.SubGroups), len(ctx.HManager.Hosts) - 1, len(ctx.HManager.Hosts))

		if err != nil {
			return err
		}
		ctx.SGManager.SubGroups = append(ctx.SGManager.SubGroups, newSg)

	} else {
		// Find the subgroup which has the minimum of hosts and let the chunk join it
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