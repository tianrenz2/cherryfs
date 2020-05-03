package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	pb "cherryfs/pkg/meta/serverpb"
	"cherryfs/pkg/meta/allocator"
	"cherryfs/pkg/object"
)

const (
	port = ":50051"
	addr = "localhost:50051"
	defaultName = "metaServer"
)

type server struct {
	pb.UnimplementedAskPutServer
}

func StartServer()  {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAskPutServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func AskPut(askPutReq *pb.AskPutRequest) ([]allocator.Target, error) {
	objName := *askPutReq.Name
	objSize := *askPutReq.Size
	objHash := *askPutReq.ObjectHash

	obj := object.Object{Name:objName, Size:int64(objSize), Hash:objHash}

	alloc := allocator.Allocator{Policy: allocator.ReplicaPolicy}

	var targets []allocator.Target
	targets, err := alloc.AllocTargets(obj)
	if err != nil {
		return targets, err
	}

	return targets, nil
}
