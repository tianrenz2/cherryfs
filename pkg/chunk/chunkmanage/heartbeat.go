package chunkmanage

import (
	"fmt"
	"log"
	"errors"
	"cherryfs/pkg/meta/watchservice"
	"cherryfs/pkg/context"
)


func StartHeartbeat() (error) {
	// Chunk services maintain their heartbeat maintaining a etcd key for each of them
	heartbeatKey := fmt.Sprintf("%s/%s", watchservice.HeartbeatPrefix, context.GlobalChunkCtx.HostId)

	ch, err := context.GlobalChunkCtx.EtcdCli.MaintainKey(heartbeatKey, watchservice.HeartbeatDeadline)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for {
		select {
			case <- context.GlobalChunkCtx.EtcdCli.Client.Ctx().Done():
				return errors.New("server closed")
			case _, ok := <-ch:
				if !ok {
					return nil
				}
		}
	}
}
