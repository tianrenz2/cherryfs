package main

import (
	pb "cherryfs/pkg/chunk/chunkserverpb"
	"google.golang.org/grpc"
	"net"
	"log"
	"fmt"
	"bytes"
	"io"
	"cherryfs/pkg/object"
	"cherryfs/pkg/chunk"
	"os"
)

type server struct {
	pb.UnimplementedPutServer
}

var chunkCtx = chunk.ChunkContext{}
var address	= ""


func StartServer()  {
	port := os.Getenv("PORT")

	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	address = os.Getenv("ADDR")
	chunkCtx.Address = address + ":" +port
	fmt.Printf("Server Address: %s\n", chunkCtx.Address)

	s := grpc.NewServer()
	pb.RegisterPutServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) PutObject(stream pb.Put_PutObjectServer) (error) {
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
	fmt.Println(data.String())
	writing = false
	err = stream.SendAndClose(&pb.PutResponse{
		Message: "Success",
		Code:   0,
	})

	lcObject := object.LocalObject{Name:name, Size: 10, Hash: hash, Path:targets[0].DestDir}

	lcObject.ObjectStore(data)

	return err
}


func main()  {
	StartServer()
}