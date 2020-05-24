package main

import (
	pb "cherryfs/pkg/comm/pb"
	"time"
	"log"
	"context"
	"google.golang.org/grpc"
	"io"
	"fmt"
	"os"
)

const (
	port     = "50051"
)

type Client struct {
	ChunkClient pb.PutClient
	MetaClient pb.MetaServiceClient
	Ctx context.Context
}

func (client *Client)MakeMetaConn(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	client.MetaClient = pb.NewMetaServiceClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 2 * time.Second)
	client.Ctx = ctx
	//defer cancel()
}

func (client *Client)MakeChunkConn(addr string) {
	fmt.Printf("Making connection to %s\n", addr)
	conn, _ := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	client.ChunkClient = pb.NewPutClient(conn)
}

func (client *Client)Put(name string, objPath string, addr string) (error) {
	addr = addr + ":" + port

	client.MakeMetaConn(addr)

	var size = int64(10)
	var objectHash = "xxxxxxxx"
	var request *pb.AskPutRequest
	request = new(pb.AskPutRequest)
	*request = pb.AskPutRequest{Name: name, Size: size, ObjectHash: objectHash}

	resp, err := client.MetaClient.AskPut(client.Ctx, request)

	if err != nil {
		return err
	}

	targets := resp.Targets

	_, err = client.UploadObject(context.Background(), objPath, targets, name, objectHash)

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	return nil
}

func (client *Client)UploadObject(ctx context.Context, f string, targets []*pb.Target, name, hash string) (pb.PutResponse, error) {
	writing := true
	file, err := os.Open(f)

	objName := file.Name()

	client.MakeChunkConn(targets[0].DestAddr)
	stream, err := client.ChunkClient.PutObject(ctx)

	sendTargets := make([]*pb.Target, 0)

	for _, target := range targets {
		sendTargets = append(
			sendTargets,
			&pb.Target{
				DestDir: target.DestDir,
				DestAddr: target.DestAddr,
			},
		)
	}

	objInfo := pb.ObjectInfo{Targets:sendTargets, Name:objName, Hash: hash}

	err = stream.Send(
		&pb.PutRequest{
			Data: &pb.PutRequest_Info{Info: &objInfo},
		})

	if err != nil {
		return pb.PutResponse{}, err
	}
	buf := make([]byte, 1024)
	for writing {
		n, err := file.Read(buf)

		if err != nil {
			if err == io.EOF {
				writing = false
				err = nil
				continue
			}
			err = fmt.Errorf("errored while copying from file to buf")
		}
		err = stream.Send(
			&pb.PutRequest{
				Data: &pb.PutRequest_Content{
					Content: buf[:n],
				},
			},
		)

		if err != nil {
			err = fmt.Errorf("failed to send chunk via stream: %v", err)
			return pb.PutResponse{}, err
		}
	}

	ret, e := stream.CloseAndRecv()

	if e != nil {
		return pb.PutResponse{}, e
	}

	return *ret, err
}


func initConnection() (pb.MetaServiceClient, context.Context) {
	var address = "127.0.0.1:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	c := pb.NewMetaServiceClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	//defer cancel()
	return c, ctx
}

func main()  {
	cli := Client{}

	if len(os.Args) < 3 {
		log.Fatalf("Not enough arguments")
		return
	}

	ObjectKey := os.Args[1]
	ObjectPath := os.Args[2]

	err := cli.Put(ObjectKey, ObjectPath, "127.0.0.1")

	if err != nil {
		fmt.Errorf("failed to put object: %v\n", err)
	}
}
