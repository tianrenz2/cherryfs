package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	pb "cherryfs/pkg/meta/server/serverpb"
	"cherryfs/pkg/meta/allocator"
	"cherryfs/pkg/object"
	cherryCtx "cherryfs/pkg/context"
	"context"
	"cherryfs/pkg/meta"
)

const (
	port = ":50051"
	addr = "localhost:50051"
	defaultName = "metaServer"
)

type server struct {
	pb.UnimplementedAskPutServer
}

var GlobalCtx cherryCtx.Context

func StartServer()  {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	GlobalCtx = meta.LoadClusterConfig()

	s := grpc.NewServer()
	pb.RegisterAskPutServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server)AskPut(ctx context.Context, askPutReq *pb.AskPutRequest) (*pb.AskPutResponse, error) {
	objName := *askPutReq.Name
	objSize := *askPutReq.Size
	objHash := *askPutReq.ObjectHash

	obj := object.Object{Name:objName, Size:int64(objSize), Hash:objHash}

	alloc := allocator.Allocator{Policy: allocator.ReplicaPolicy, Ctx: GlobalCtx}

	targets, err := alloc.AllocTargets(obj)

	if err != nil {
		return nil, err
	}

	var status = int32(0)

	respTargets := make([]*pb.Habitat, 0)
	for _, target := range targets {
		respTargets = append(respTargets, &pb.Habitat{DestAddr: &target.Host.Address, DestDir:&target.Dir.Path})
	}

	var resp = pb.AskPutResponse{Status:&status, Targets:respTargets}

	return &resp, nil
}

func main()  {
	StartServer()
}