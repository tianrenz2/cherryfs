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
	pb.UnimplementedPutServer
}

func startServer()  {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPutServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main()  {
	startServer()
}