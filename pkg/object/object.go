package object

import (
	"cherryfs/pkg/comm/pb"
	//"cherryfs/pkg/meta/server"
	"encoding/json"
	"cherryfs/pkg/context"
	"fmt"
	"os"
)

type Object struct {
	Name string
	Size int64
	Hash string
	Targets []*pb.Target
}

const (
	ObjectKeyPrefix = "Object/"
)

func (Obj *Object) PutMeta(ctx context.Context) (error) {
	serialized, err := json.Marshal(Obj)

	if err != nil {
		return err
	}

	key := ObjectKeyPrefix + Obj.Name

	fmt.Println(key)
	ctx.EtcdCli.CreateEtcdClient(os.Getenv("ETCDADDR"))
	err = ctx.EtcdCli.Put(key, string(serialized))

	for _, target := range Obj.Targets {
		ctx.EtcdCli.Put(ObjectKeyPrefix + Obj.Name + "/" + target.DestId, "0")
	}

	if err != nil {
		return err
	}

	return nil
}

func GetObjectTarget(name string, ctx context.Context) (*pb.Target, error) {
	ctx.EtcdCli.CreateEtcdClient(os.Getenv("ETCDADDR"))

	info, err := ctx.EtcdCli.Get(ObjectKeyPrefix + name)

	if err != nil {
		return nil, fmt.Errorf("failed to get object: %v", err)
	}

	var object Object

	err = json.Unmarshal([]byte(info), &object)

	if err != nil {
		return nil, fmt.Errorf("failed to parse the object: %v", err)
	}

	state := "0"
	chosenIndex := 0

	for {
		getKey := ObjectKeyPrefix + "/" + name + "/" + object.Targets[chosenIndex].DestId

		fmt.Printf("Get by key: %s\n", getKey)
		state, _ = ctx.EtcdCli.Get(getKey)
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

