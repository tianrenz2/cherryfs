package chunkmanage

import (
	"fmt"
	"log"
	"errors"
)

const (
	HeartbeatPrefix = "heartbeat"
	HeartbeatDeadline = 10
)


func (chunkCtx *ChunkContext) StartHeartbeat() (error) {
	// Chunk services maintain their heartbeat maintaining a etcd key for each of them
	heartbeatKey := fmt.Sprintf("%s/%s", HeartbeatPrefix, chunkCtx.HostId)

	ch, err := chunkCtx.EtcdCli.MaintainKey(heartbeatKey, HeartbeatDeadline)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for {
		select {
			case <- chunkCtx.EtcdCli.Client.Ctx().Done():
				return errors.New("server closed")
			case ka, ok := <-ch:
				if !ok {
					return nil
				} else {
					log.Printf("Recv reply from chunk: %s, ttl:%d", chunkCtx.HostId, ka.TTL)
				}
		}
	}

	for {
		sendHeartbeat()
	}
}

func sendHeartbeat()  {
	
}