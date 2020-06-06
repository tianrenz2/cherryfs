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

func StartServer()  {
	//chunkContext = initContext()
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
			err = lcObject.PostStore(chunkCtx)
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

func (s *ChunkServer) GetObject(getRequest *pb.GetRequest, responser pb.ChunkServer_GetObjectServer) (error) {
	dir := getRequest.Dir
	name := getRequest.Name

	file := dir + "/" + name
	f, _ := os.Open(file)
	buf := make([]byte, 1024)

	sending := true

	for sending {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				sending = false
				err = nil
				break
			}
			err = fmt.Errorf("errored while copying from file to buf")
		}

		err = responser.Send(
			& pb.GetResponse{
				Content: buf[:n],
			},
		)

		if err != nil {
			err = fmt.Errorf("failed to send chunkmanage via stream: %v", err)
			return err
		}
	}

	return nil
}


func initContext() chunkmanage.ChunkContext {
	var etcdClient etcd.EtcdClient
	etcdClient.CreateEtcdClient(os.Getenv("ETCDADDR"))

	newCtx := chunkmanage.ChunkContext{
		EtcdCli: etcdClient,
	}

	return newCtx
}


func main()  {
	var etcdClient etcd.EtcdClient
	etcdClient.CreateEtcdClient(os.Getenv("ETCDADDR"))
	chunkCtx.EtcdCli = etcdClient

	address = os.Getenv("ADDR")
	port := os.Getenv("PORT")
	chunkCtx.Address = address + ":" + port

	chunkCtx.StartupChunk()
	go chunkCtx.StartHeartbeat()

	StartServer()
}

func testPut() {
	lcobj := object.LocalObject{Name: "abc"}
	err := lcobj.PostStore(chunkCtx)
	fmt.Println(err)
}