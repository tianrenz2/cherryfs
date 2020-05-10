package main

import (
	pb "cherryfs/pkg/meta/server/serverpb"
	"time"
	"log"
	"context"
	"google.golang.org/grpc"
	"cherryfs/pkg/chunk/chunkserverpb"
)

const (
	address     = "localhost:50051"
)

type client struct {
	chunkserverpb.PutClient
}

func initConnection() (pb.AskPutClient, context.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAskPutClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return c, ctx
}

func main()  {

	var objName = "obj1"
	var size = int64(100)
	var object = "asdfasdfasdfasdfasdf"

	var objectHash = "asdfasdfgergergergefgdfgdfgdfg"

	c, ctx := initConnection()

	r, err := c.AskPut(ctx, &pb.AskPutRequest{Name: &objName, Size: &size, Object: &object, ObjectHash: &objectHash})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", r.Targets)
}
