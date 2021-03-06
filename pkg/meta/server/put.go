package server

import (
	"cherryfs/pkg/comm/pb"
	"cherryfs/pkg/object"
	"cherryfs/pkg/meta/allocator"
	"context"
	"fmt"
	context2 "cherryfs/pkg/context"
)


func (s *MetaServer)AskPut(ctx context.Context, askPutReq *pb.AskPutRequest) (*pb.AskPutResponse, error) {
	objName := askPutReq.Name
	objSize := askPutReq.Size
	objHash := askPutReq.ObjectHash

	obj := object.Object{Name:objName, Size:int64(objSize), Hash:objHash}

	alloc := allocator.Allocator{Policy: allocator.ReplicaPolicy, Ctx: *context2.GlobalCtx}

	targets, err := alloc.AllocTargets(obj)

	fmt.Println(targets)

	if err != nil {
		return nil, err
	}

	var status = int32(0)

	respTargets := make([]*pb.Target, 0)
	for _, target := range targets {
		respTargets = append(
			respTargets,
			&pb.Target{
				DestAddr: target.Host.Address,
				DestDir:target.Dir.Path,
				DestId:target.Host.HostId,
				SgId:int32(target.SgId),
			})
	}

	obj.Targets = respTargets
	obj.PutMeta(context2.GlobalCtx.EtcdCli)

	var resp = pb.AskPutResponse{Status:status, Targets:respTargets}

	return &resp, nil
}

