package object

import (
	"cherryfs/pkg/comm/pb"
	"encoding/json"
	"fmt"
	"cherryfs/internal/etcd"
)

type Object struct {
	Name string
	Size int64
	Hash string
	Targets []*pb.Target
}

const (
	ObjectKeyPrefix = "Object"
)

func (Obj *Object) PutMeta(client etcd.EtcdClient) (error) {
	serialized, err := json.Marshal(Obj)

	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s/%s", ObjectKeyPrefix, Obj.Name)

	fmt.Println(key)
	//ctx.EtcdCli.CreateEtcdClient(os.Getenv("ETCDADDR"))
	err = client.Put(key, string(serialized))

	for _, target := range Obj.Targets {
		putkey := fmt.Sprintf("%s/%s/%s", ObjectKeyPrefix, Obj.Name, target.DestId)
		client.Put(putkey, "0")
	}

	if err != nil {
		return err
	}

	return nil
}

func GetObjectTarget(name string, client etcd.EtcdClient) (*pb.Target, error) {
	//ctx.EtcdCli.CreateEtcdClient(os.Getenv("ETCDADDR"))

	objectKey := fmt.Sprintf("%s/%s", ObjectKeyPrefix, name)

	info, err := client.Get(objectKey)

	if err != nil {
		return nil, fmt.Errorf("failed to get object: %v", err)
	}

	var object Object

	err = json.Unmarshal([]byte(info), &object)

	fmt.Printf("info: %v\n", info)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the object: %v", err)
	}

	state := "0"
	chosenIndex := 0

	for {
		getKey := fmt.Sprintf("%s/%s/%s", ObjectKeyPrefix, object.Targets[chosenIndex].DestId, name)

		fmt.Printf("Get by key: %s\n", getKey)
		state, _ = client.Get(getKey)
		if state == "1" {
			break
		}
		chosenIndex += 1
		if chosenIndex >= len(object.Targets) {
			return nil, fmt.Errorf("no target is valid")
		}
	}

	return object.Targets[chosenIndex], nil
}

