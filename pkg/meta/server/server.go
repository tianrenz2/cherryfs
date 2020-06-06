package server

import (
	"cherryfs/pkg/comm/pb"
	cherryCtx "cherryfs/pkg/context"
	"cherryfs/pkg/meta/initialize"
	"log"
	"google.golang.org/grpc"
	"net"
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

	go GlobalCtx.RegistryWatcher()
	go GlobalCtx.HeartbeatWatcher()

	s := grpc.NewServer()
	pb.RegisterMetaServiceServer(s, &MetaServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main()  {
	StartServer()
}