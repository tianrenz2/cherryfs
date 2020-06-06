package context

import (
	"go.etcd.io/etcd/mvcc/mvccpb"
	"log"
)


const (
	HeartbeatPrefix = "heartbeat"
	HeartbeatDeadline = 10
)


func (ctx *Context) HeartbeatWatcher() {
	// Registry Watcher is watching for new (chunk) services for registration
	watchChan := ctx.EtcdCli.WatchKey(true, HeartbeatPrefix)

	for res := range watchChan {
		evType := res.Events[0].Type
		switch evType {
			case mvccpb.PUT:
				log.Printf("%s got put\n", string(res.Events[0].Kv.Key))

			case mvccpb.DELETE:
				log.Printf("%s got deleted\n", string(res.Events[0].Kv.Key))
		}
	}
}
