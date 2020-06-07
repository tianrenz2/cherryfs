package main

import (
	"cherryfs/pkg/comm/pb"
	"time"
	"log"
	"context"
	"google.golang.org/grpc"
	"fmt"
	"os"
	"io"
	"bytes"
	"io/ioutil"
)

const (
	port     = "50051"
)

type Client struct {
	ChunkClient pb.ChunkServerClient
	MetaClient pb.MetaServiceClient
	Ctx context.Context
	ClusterAddr string
}

func (client *Client)Init(clusterAddr string)  {
	client.ClusterAddr = clusterAddr
}

func (client *Client)MakeMetaConn() {
	conn, err := grpc.Dial(client.ClusterAddr, grpc.WithInsecure(), grpc.WithBlock())
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
	client.ChunkClient = pb.NewChunkServerClient(conn)
}

func (client *Client)Put(name string, objPath string) (error) {
	client.MakeMetaConn()

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

	for _, t := range targets{
		fmt.Printf("%s, %s \n", t.DestAddr, t.DestDir)
	}


	_, err = client.UploadObject(context.Background(), objPath, targets, name, objectHash)

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	return nil
}

func (client *Client)UploadObject(ctx context.Context, f string, targets []*pb.Target, name, hash string) (pb.PutResponse, error) {
	writing := true
	file, err := os.Open(f)

	objName := name

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
			err = fmt.Errorf("failed to send chunkmanage via stream: %v", err)
			return pb.PutResponse{}, err
		}
	}

	ret, e := stream.CloseAndRecv()

	if e != nil {
		return pb.PutResponse{}, e
	}

	return *ret, err
}

func (client *Client) Get(name, outputFile string) (error) {
	client.MakeMetaConn()

	askGetReq := pb.AskGetRequest{Name:name}

	askGetResp, err := client.MetaClient.AskGet(client.Ctx, &askGetReq)

	if err != nil {
		return fmt.Errorf("failed to get object: %v", err)
	}

	target := askGetResp.Target
	fmt.Println(target)
	client.DownloadObject(target.DestAddr, target.DestDir, name, outputFile)

	return nil
}

func (client *Client) DownloadObject(addr, dir, name, outputFile string) (error) {
	client.MakeChunkConn(addr)
	stream, err := client.ChunkClient.GetObject(context.Background(), &pb.GetRequest{Name: name, Dir: dir})

	if err != nil {
		return err
	}
	os.Create(outputFile)

	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}

	data := bytes.Buffer{}
	for {
		recChunk, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			err = fmt.Errorf("failed unexpectadely while reading chunks from stream")
			return err
		}else {
			data.Write(recChunk.GetContent())
		}
	}

	ioutil.WriteFile(outputFile, data.Bytes(), 0644)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

func main()  {
	cli := Client{}

	cli.Init("127.0.0.1:50051")

	if len(os.Args) < 3 {
		log.Fatalf("Not enough arguments")
		return
	}

	ObjectKey := os.Args[1]
	ObjectPath := os.Args[2]

	err := cli.Put(ObjectKey, ObjectPath)
	//err := cli.Get(ObjectKey, ObjectPath)

	if err != nil {
		fmt.Errorf("failed to put object: %v\n", err)
	}
}
