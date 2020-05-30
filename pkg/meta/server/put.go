package server

import (
	"cherryfs/pkg/comm/pb"
	"cherryfs/pkg/object"
	"cherryfs/pkg/meta/allocator"
	"context"
)


func (s *MetaServer)AskPut(ctx context.Context, askPutReq *pb.AskPutRequest) (*pb.AskPutResponse, error) {
	objName := askPutReq.Name
	objSize := askPutReq.Size
	objHash := askPutReq.ObjectHash

	obj := object.Object{Name:objName, Size:int64(objSize), Hash:objHash}

	alloc := allocator.Allocator{Policy: allocator.ReplicaPolicy, Ctx: GlobalCtx}

	targets, err := alloc.AllocTargets(obj)

	if err != nil {
		return nil, err
	}

	var status = int32(0)

	respTargets := make([]*pb.Target, 0)
	for _, target := range targets {
		respTargets = append(respTargets, &pb.Target{DestAddr: target.Host.Address, DestDir:target.Dir.Path, DestId:target.Host.HostId})
	}

	obj.Targets = respTargets
	obj.PutMeta(GlobalCtx)

	var resp = pb.AskPutResponse{Status:status, Targets:respTargets}

	return &resp, nil
}

