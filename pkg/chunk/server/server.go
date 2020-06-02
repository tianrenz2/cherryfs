package main

import (
	"cherryfs/pkg/comm/pb"
	"google.golang.org/grpc"
	"net"
	"log"
	"fmt"
	"bytes"
	"io"
	"cherryfs/pkg/object"
	"cherryfs/pkg/chunk/chunkmanage"
	"os"
	"cherryfs/pkg/etcd"
)

type ChunkServer struct {
	pb.UnimplementedChunkServerServer
}

var chunkCtx = chunkmanage.ChunkContext{}
var address	= ""


var chunkContext chunkmanage.ChunkContext

func StartServer()  {
	chunkContext = initContext()

	address = os.Getenv("ADDR")
	port := os.Getenv("PORT")
	chunkCtx.Address = address + ":" + port

	lis, err := net.Listen("tcp", chunkCtx.Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Printf("Server Address: %s\n", chunkCtx.Address)

	s := grpc.NewServer()
	pb.RegisterChunkServerServer(s, &ChunkServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}


func (s *ChunkServer) PutObject(stream pb.ChunkServer_PutObjectServer) (error) {
	info, err := stream.Recv()
	if err != nil {
		return err
	}

	targets := info.GetInfo().Targets
	name := info.GetInfo().Name
	hash := info.GetInfo().Hash

	fmt.Printf("Received: %s\n", name)

	data := bytes.Buffer{}

	var writing = true
	go chunkCtx.SendAsRec(&writing, &data, info.GetInfo(), targets)
	for {
		recChunk, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			}
			err = fmt.Errorf("failed unexpectadely while reading chunks from stream")
			writing = false
			return err
		}else {
			data.Write(recChunk.GetContent())
		}
	}
	writing = false

	var selfTarget *pb.Target
	for _, target := range targets {
		if target.DestAddr == chunkCtx.Address {
			selfTarget = target
		}
	}

	if selfTarget != nil {
		fmt.Printf("Writing to %s\n", selfTarget.DestDir)
		lcObject := object.LocalObject{Name:name, Size: 10, Hash: hash, Path:selfTarget.DestDir}
		err = lcObject.ObjectStore(data)
		if err == nil{
			err = lcObject.PostStore(chunkContext)
		} else {
			fmt.Printf("error happened %v\n", err)
		}
	}
	var msg = "Success"
	var code = 0

	if err != nil {
		msg = err.Error()
		code = 1
	}

	err = stream.SendAndClose(&pb.PutResponse{
		Message: msg,
		Code:   int32(code),
	})

	return err
}

func initContext() chunkmanage.ChunkContext {
	var etcdClient etcd.EtcdClient
	etcdClient.CreateEtcdClient(os.Getenv("ETCD_ADDR"))

	newCtx := chunkmanage.ChunkContext{
		EtcdCli: etcdClient,
	}

	return newCtx
}


func main()  {
	var etcdClient etcd.EtcdClient
	etcdClient.CreateEtcdClient(os.Getenv("ETCDADDR"))
	cCtx := chunkmanage.ChunkContext{
		EtcdCli:etcdClient,
		HostId: "xxxxxx",
	}

	cCtx.StartHeartbeat()

	//cCtx.StartupChunk()
	StartServer()
}

func testPut() {
	lcobj := object.LocalObject{Name: "abc"}
	err := lcobj.PostStore(chunkContext)
	fmt.Println(err)
}