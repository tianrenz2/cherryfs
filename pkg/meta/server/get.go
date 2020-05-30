package server

import (
	"cherryfs/pkg/comm/pb"
	"context"
	"cherryfs/pkg/object"
)


func (s *MetaServer)AskGet(ctx context.Context, askGetReq *pb.AskGetRequest) (*pb.AskGetResponse, error) {
	name := askGetReq.Name

	objTarget, err := object.GetObjectTarget(name, GlobalCtx)

	if err != nil {
		return nil, err
	}

	resp := pb.AskGetResponse{Target: objTarget}

	return &resp, nil
}
