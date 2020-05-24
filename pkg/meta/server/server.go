package server

import (
	"net"
	"log"
	"google.golang.org/grpc"
	pb "cherryfs/pkg/comm/pb"
	"cherryfs/pkg/meta/allocator"
	"cherryfs/pkg/object"
	cherryCtx "cherryfs/pkg/context"
	"context"
	"cherryfs/pkg/meta/initialize"
	"fmt"
)

const (
	port = ":50051"
	addr = "localhost:50051"
	defaultName = "metaServer"
)

type MetaServer struct {
	pb.MetaServiceServer
}

var GlobalCtx cherryCtx.Context

func StartServer()  {
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	GlobalCtx = initialize.Startup()

	fmt.Println("\n" + lis.Addr().String())

	s := grpc.NewServer()
	pb.RegisterMetaServiceServer(s, &MetaServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

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

	var resp = pb.AskPutResponse{Status:status, Targets:respTargets}

	return &resp, nil
}

func main()  {
	StartServer()
}