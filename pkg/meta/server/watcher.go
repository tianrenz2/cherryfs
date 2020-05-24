package server
//
//import (
//	"cherryfs/pkg/comm/heartbeatpb"
//	"net"
//	"cherryfs/pkg/meta/initialize"
//	"google.golang.org/grpc"
//	"log"
//	"cherryfs/pkg/context"
//	"cherryfs/pkg/etcd"
//)
//
//const (
//	HbPort = ":50052"
//	HbAddr = "localhost:50052"
//)
//
//type hbServer struct {
//	heartbeatpb.UnimplementedRecHeartBeatServer
//}
//
//func StartHeartbeatServer()  {
//	lis, err := net.Listen("tcp", port)
//	if err != nil {
//		log.Fatalf("Failed to listen: %v", err)
//	}
//
//	GlobalCtx = initialize.LoadClusterConfig()
//
//	s := grpc.NewServer()
//	heartbeatpb.RegisterRecHeartBeatServer(s, &hbServer{})
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("Failed to serve: %v", err)
//	}
//}
//
//func (*hbServer)RecHeartBeat(ctx context.Context, request *heartbeatpb.HeartbeatRequest) (*heartbeatpb.HeartbeatResponse, error) {
//	etcdCli := etcd.EtcdClient{}
//
//	return nil, nil
//}