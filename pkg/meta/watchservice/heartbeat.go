package watchservice

import (
	"github.com/coreos/etcd/mvcc/mvccpb"
	"log"
	"strings"
	"cherryfs/pkg/context"
)


const (
	HeartbeatPrefix = "heartbeat"
	HeartbeatDeadline = 1
)


func HeartbeatWatcher() {
	EtcdCli := context.GlobalCtx.EtcdCli
	// Registry Watcher is watching for new (chunk) services for registration
	watchChan := EtcdCli.WatchKey(true, HeartbeatPrefix)

	for res := range watchChan {
		evType := res.Events[0].Type
		switch evType {
			case mvccpb.PUT:
				log.Printf("%s got put\n", string(res.Events[0].Kv.Key))

			case mvccpb.DELETE:
				deletedKey := string(res.Events[0].Kv.Key)
				lostHostId := strings.Split(deletedKey, "/")[1]
				log.Printf("%s got lost\n", lostHostId)
				go LostHostHandler(lostHostId)
		}
	}
}

func LostHostHandler(hostId string) {
	lostHost, err := context.GlobalCtx.HManager.GetHostPointerByHostId(hostId)

	if err != nil {
		log.Fatalf("failed to get the host %v \n", err)
	}
	lostHost.ClaimAsLost()
	log.Printf("claimed host %s \n", lostHost.HostId)

	log.Printf("ctx: %v\n", context.GlobalCtx)
	recoverDestSgId := lostHost.SubgroupId

	recoverSrcSgId := recoverDestSgId + 1 % (len(context.GlobalCtx.SGManager.SubGroups))
	objects, _ := lostHost.GetAllObjects(context.GlobalCtx.EtcdCli)

	recoverMap, err := GenerateRecoverMap(context.GlobalCtx, recoverSrcSgId, recoverDestSgId, objects)

	err = SetupRecoverTransmissions(recoverMap)

	if err != nil {
		log.Printf("%v", err)
	}

}

