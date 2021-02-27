package main

import (
	"cherryfs/pkg/comm/pb"
	cherryCtx "cherryfs/pkg/context"
	"log"
	"google.golang.org/grpc"
	"net"
	"cherryfs/pkg/meta/watchservice"
	"cherryfs/pkg/meta/initialize"
)

const (
	port = ":50051"
	addr = "localhost:50051"
	defaultName = "metaServer"
)

type MetaServer struct {
	pb.MetaServiceServer
}

func StartServer()  {
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	cherryCtx.GlobalCtx = new(cherryCtx.Context)
	*cherryCtx.GlobalCtx = initialize.LoadClusterConfig()

	log.Printf("global ctx: %v\n", cherryCtx.GlobalCtx.EtcdCli)

	log.Printf("etcd client: %v\n", cherryCtx.GlobalCtx.EtcdCli)

	go cherryCtx.GlobalCtx.RegistryWatcher()
	go watchservice.HeartbeatWatcher()

	s := grpc.NewServer()
	pb.RegisterMetaServiceServer(s, &MetaServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main()  {
	StartServer()
}