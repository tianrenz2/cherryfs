package main

import (
	"cherryfs/pkg/comm/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"os"
)

type PutClientIns struct {
	client pb.PutClient
}

func (c *PutClientIns) UploadObject(ctx context.Context, f string) (pb.PutResponse, error) {

	objName := "objname1"
	writing := true
	file, err := os.Open(f)

	stream, err := c.client.PutObject(ctx)
	targets := []*pb.Target{
		{DestDir:"/tmp/cherryfs3", DestAddr:"192.168.30.2:5001"},
	}

	for i := 0; i < 3; i++ {
		targets = append(targets, )
	}

	objInfo := pb.ObjectInfo{Targets:targets, Name:objName, Hash: "randomhash"}

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
		fmt.Println(string(buf[:n]))
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

func main()  {
	var client = PutClientIns{}

	req_addr := os.Args[1]

	conn, _ := grpc.Dial(req_addr, grpc.WithInsecure(), grpc.WithBlock())
	client.client = pb.NewPutClient(conn)

	file := "t.txt"

	status, e := client.UploadObject(context.Background(), file)

	if e != nil {
		fmt.Printf("error: %v", e)
	}

	fmt.Printf("Status: %d, %s\n", status.Code, status.Message)
}
