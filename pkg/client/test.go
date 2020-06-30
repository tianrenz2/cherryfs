package client

import (
	pb "cherryfs/pkg/comm/pb/test"
	"google.golang.org/grpc"
	"time"
	"context"
	"log"
	"fmt"
)

func main()  {
	c, ctx := initTestConnection()
	var name = "aaaaa"
	r, err := c.TestPut(ctx, &pb.TestPutRequest{Name: &name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", r.Status)

}

func initTestConnection() (pb.TestPutServiceClient, context.Context) {
	var address = "127.0.0.1:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()

	fmt.Println(conn.GetState())
	c := pb.NewTestPutServiceClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	//defer cancel()
	return c, ctx
}
