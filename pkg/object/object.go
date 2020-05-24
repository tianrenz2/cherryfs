package object

import (
	"cherryfs/pkg/comm/pb"
	//"cherryfs/pkg/meta/server"
	"encoding/json"
	"cherryfs/pkg/context"
	"fmt"
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

	err = ctx.EtcdCli.Put(ObjectKeyPrefix + Obj.Name, string(serialized))

	for _, target := range Obj.Targets {
		ctx.EtcdCli.Put(ObjectKeyPrefix + Obj.Name + "/" + target.DestId, "0")
	}

	if err != nil {
		return err
	}

	return nil
}

func GetObject(name string, ctx context.Context) (*pb.Target, error) {

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

	for state=="0" {
		state, _ = ctx.EtcdCli.Get(ObjectKeyPrefix + name + object.Targets[chosenIndex].DestId)
		chosenIndex += 1
		if chosenIndex >= len(object.Targets) {
			return nil, fmt.Errorf("no target is valid")
		}
	}

	return object.Targets[chosenIndex], nil
}