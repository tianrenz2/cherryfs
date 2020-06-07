package watchservice

import (
	"go.etcd.io/etcd/mvcc/mvccpb"
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

	//context.GlobalChunkCtx.EtcdCli = etcd.EtcdClient{}
	//context.GlobalChunkCtx.EtcdCli.CreateEtcdClient(os.Getenv("ETCDADDR"))
	//
	objects, _ := lostHost.GetAllObjects(context.GlobalCtx.EtcdCli)
	//
	for _, obj := range objects {
		log.Printf("objects to recover: %v", obj.Name)
	}

}