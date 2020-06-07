package main

import (
	"cherryfs/pkg/object"
	"fmt"
	"cherryfs/pkg/comm/pb"
	"os"
	"io"
	"google.golang.org/grpc"
	"time"
	"context"
	"path"
	"log"
)

type TaskState int8
type TaskType int8

const (
	PROCESSING TaskState = 0
	SUCCEEDED TaskState = 1
	FAILED TaskState = 2

	CopyObjects TaskType = 1
)

type Task struct {
	RunnerHostId string
	State TaskState
	Type TaskType
	Info interface{}
}

type CopyTaskInfo struct {
	Target pb.Target
	LocalObjects []object.LocalObject
}

func (TaskExecutor *Task) ExecuteTask() (err error) {
	switch TaskExecutor.Type {
		case CopyObjects:
			TaskExecutor.CopyObjectsTask()
	}

	return nil
}

func (TaskExecutor *Task) CopyObjectsTask() (err error) {
	copyMap := TaskExecutor.Info.([]CopyTaskInfo)

	log.Printf("start recovering objects \n")
	for _, copyTaskInfo := range copyMap {

		target := pb.Target{
			DestId: copyTaskInfo.Target.DestId,
			DestAddr: copyTaskInfo.Target.DestAddr,
			DestDir: copyTaskInfo.Target.DestDir,
		}

		TaskExecutor.CopyObjects(target, copyTaskInfo.LocalObjects)
	}

	return nil
}

func (TaskExecutor *Task)CopyObjects(target pb.Target, objectList []object.LocalObject) (pb.PutResponse, error) {
	conn, _ := grpc.Dial(target.DestAddr, grpc.WithInsecure(), grpc.WithBlock())
	chunkClient := pb.NewChunkServerClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	stream, err := chunkClient.CopyObject(ctx)

	sendTargets := make([]*pb.Target, 0)
	sendTargets = append(sendTargets, &target)


	for _, obj := range objectList {
		objInfo := pb.ObjectInfo{
			Targets: sendTargets, Name: obj.Name, Hash: obj.Hash,
		}

		file, err := os.Open(path.Join(obj.Path, obj.Name))

		if err != nil {
			log.Fatalf("failed to read object: %v\n", err)
		}

		err = stream.Send(
			&pb.PutRequest{
				Data: &pb.PutRequest_Info{Info: &objInfo},
			})

		if err != nil {
			return pb.PutResponse{}, err
		}
		buf := make([]byte, 1024)
		writing := true
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
	}

	ret, e := stream.CloseAndRecv()

	if e != nil {
		return pb.PutResponse{}, e
	}

	return *ret, err
}

func main()  {
	target := pb.Target{
		DestAddr: "192.168.30.4:50011",
		DestDir: "/tmp/cherryfs/2/",
		DestId:"",
	}

	lcObject := object.LocalObject{
		Name: "aaa",
		Path: "/tmp/cherryfs/1/",

	}

	lcObjectList := make([]object.LocalObject, 0)
	lcObjectList = append(lcObjectList, lcObject)

	taskExecutor := Task{}
	taskExecutor.CopyObjects(target, lcObjectList)
}