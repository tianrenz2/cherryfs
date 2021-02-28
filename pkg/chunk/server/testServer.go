package server

import (
	"net"
	"google.golang.org/grpc"
	pb "cherryfs/pkg/comm/pb/test"
	"log"
	"context"
	"fmt"
)

type TestServer struct {
	pb.UnimplementedTestPutServiceServer
}


func StartTestServer()  {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Listen " + lis.Addr().String())
	//fmt.Println(lis.Accept())
	s := grpc.NewServer()

	pb.RegisterTestPutServiceServer(s, &TestServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *TestServer)TestPut(ctx context.Context, in *pb.TestPutRequest) (*pb.TestPutResponse, error) {
	fmt.Println("received")

	var status = int32(1)
	resp := pb.TestPutResponse{
		Status: &status,
	}

	return &resp, nil
}