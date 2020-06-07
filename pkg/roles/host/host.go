package host

import (
	"cherryfs/pkg/etcd"
	"fmt"
	"cherryfs/pkg/object"
	"strings"
	"log"
	"encoding/json"
	"os"
)

type HostState int32

const (
	HEALTHY HostState = 0
	OFFLINE HostState = 1
)

type ConfigHost struct {
	Hostname string
	Address string
	Dirs []string
}

type Host struct {
	SubgroupId int
	HostId string
	Hostname string
	Address string
	Dirs []string
	HostState HostState
}

func (Host *Host) ClaimAsLost() {
	Host.HostState = OFFLINE
}

func (Host *Host) GetAllObjects(etcdClient etcd.EtcdClient) ([]object.Object, error) {
	etcdClient.CreateEtcdClient(os.Getenv("ETCDADDR"))
	prefix := fmt.Sprintf("%s/%s", object.ObjectKeyPrefix, Host.HostId)
	kvMap, err := etcdClient.GetWithPrefix(prefix)

	log.Printf("kvmap %v\n", kvMap)
	if err != nil {
		return nil, err
	}

	objects := make([]object.Object, 0)

	for key, _ := range kvMap {
		slices := strings.Split(key, "/")
		objectName := slices[len(slices) - 1]
		log.Printf("to recover object %s\n", objectName)

		objectKey := fmt.Sprintf("%s/%s", object.ObjectKeyPrefix, objectName)
		objInfo, err := etcdClient.Get(objectKey)

		if err != nil {
			log.Printf("failed to get information of object %s\n", objectKey)
		}

		var objectInstance object.Object
		err = json.Unmarshal([]byte(objInfo), &objectInstance)
		log.Printf("info: %v\n", objectInstance)
		objects = append(objects, objectInstance)
	}

	return objects, nil
}