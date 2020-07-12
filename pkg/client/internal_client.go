package client

import (
	"cherryfs/pkg/comm/pb"
	"google.golang.org/grpc"
	"time"
	"log"
	"context"
	"cherryfs/pkg/task"
	"encoding/json"
)


type InternalClient struct {
	ChunkClient pb.ChunkServerClient
	serverAddr string
}

func (internalClient *InternalClient) New(addr string) {
	internalClient.serverAddr = addr
	return
}

func (internalClient *InternalClient) SendTask(taskType task.TaskType, taskInfo interface{}) error {
	conn, err := grpc.Dial(internalClient.serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	internalClient.ChunkClient = pb.NewChunkServerClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 2 * time.Second)
	var taskInfoBytes []byte

	switch taskType {
	case task.CopyObjects:
		taskInfoBytes, err = json.Marshal(taskInfo.([]task.CopyTaskInfo))
		if err != nil {
			return err
		}
	}

	taskReq := pb.TaskRequest{
		TaskType: int32(taskType),
		Value: taskInfoBytes,
	}

	_, err = internalClient.ChunkClient.TaskReceiver(ctx, &taskReq)

	if err != nil {
		return err
	}

	return nil
}