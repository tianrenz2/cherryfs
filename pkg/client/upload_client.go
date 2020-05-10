package main

import (
	"cherryfs/pkg/chunk/chunkserverpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"os"
)

type PutClientIns struct {
	client chunkserverpb.PutClient
}

func (c *PutClientIns) UploadObject(ctx context.Context, f string) (chunkserverpb.PutResponse, error) {

	objName := "objname1"
	writing := true
	file, err := os.Open(f)

	stream, err := c.client.PutObject(ctx)
	targets := []*chunkserverpb.Target{
		{DestDir:"/tmp/cherryfs1", DestAddr:"192.168.30.3:50010"},
		{DestDir:"/tmp/cherryfs1", DestAddr:"192.168.30.4:50011"},
	}

	for i := 0; i < 3; i++ {
		targets = append(targets, )
	}

	objInfo := chunkserverpb.ObjectInfo{Targets:targets, Name:objName, Hash: "randomhash"}

	err = stream.Send(
		&chunkserverpb.PutRequest{
			Data: &chunkserverpb.PutRequest_Info{Info: &objInfo},
		})

	if err != nil {
		return chunkserverpb.PutResponse{}, err
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
		fmt.Println(string(buf[:n]))
		err = stream.Send(
			&chunkserverpb.PutRequest{
				Data: &chunkserverpb.PutRequest_Content{
					Content: buf[:n],
				},
			},
		)

		if err != nil {
			err = fmt.Errorf("failed to send chunk via stream: %v", err)
			return chunkserverpb.PutResponse{}, err
		}
	}

	ret, e := stream.CloseAndRecv()

	if e != nil {
		return chunkserverpb.PutResponse{}, e
	}

	return *ret, err
}

func main()  {
	var client = PutClientIns{}

	req_addr := os.Args[1]

	conn, _ := grpc.Dial(req_addr, grpc.WithInsecure(), grpc.WithBlock())
	client.client = chunkserverpb.NewPutClient(conn)

	file := "t.txt"

	status, e := client.UploadObject(context.Background(), file)

	if e != nil {
		fmt.Printf("error: %v", e)
	}

	fmt.Printf("Status: %d, %s\n", status.Code, status.Message)
}
