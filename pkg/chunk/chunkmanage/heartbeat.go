package chunkmanage

import (
	"fmt"
	"log"
	"errors"
	"cherryfs/pkg/context"
)


func (chunkCtx *ChunkContext) StartHeartbeat() (error) {
	// Chunk services maintain their heartbeat maintaining a etcd key for each of them
	heartbeatKey := fmt.Sprintf("%s/%s", context.HeartbeatPrefix, chunkCtx.HostId)

	ch, err := chunkCtx.EtcdCli.MaintainKey(heartbeatKey, context.HeartbeatDeadline)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for {
		select {
			case <- chunkCtx.EtcdCli.Client.Ctx().Done():
				return errors.New("server closed")
			case _, ok := <-ch:
				if !ok {
					return nil
				}
		}
	}

	for {
		sendHeartbeat()
	}
}

func sendHeartbeat()  {
	
}