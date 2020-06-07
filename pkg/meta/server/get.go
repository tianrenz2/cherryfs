package server

import (
	"cherryfs/pkg/comm/pb"
	"context"
	"cherryfs/pkg/object"
	context2 "cherryfs/pkg/context"
)


func (s *MetaServer)AskGet(ctx context.Context, askGetReq *pb.AskGetRequest) (*pb.AskGetResponse, error) {
	name := askGetReq.Name

	objTarget, err := object.GetObjectTarget(name, context2.GlobalCtx.EtcdCli)

	if err != nil {
		return nil, err
	}

	resp := pb.AskGetResponse{Target: objTarget}

	return &resp, nil
}
