package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	pb "cherryfs/pkg/meta/serverpb"
)

const (
	port = ":50051"
	addr = "localhost:50051"
	defaultName = "metaServer"
)

type server struct {
	pb.UnimplementedAskPutServer
}

func startServer()  {
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

func AskPut(askPutReq pb.AskPutRequest) {
	objName := askPutReq.Name
	objHash := askPutReq.ObjectHash


}

func main()  {
	startServer()
}