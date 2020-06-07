package chunkmanage

import (
	"bytes"
	"cherryfs/pkg/comm/pb"
	"google.golang.org/grpc"
	"fmt"
	"context"
	context2 "cherryfs/pkg/context"
)

func SendAsRec(writing *bool, buffer *bytes.Buffer, info *pb.ObjectInfo, targets []*pb.Target) (error) {
	var newTargets = make([]*pb.Target, 0)
	for _, target := range targets {
		if target.DestAddr == context2.GlobalChunkCtx.Address {
			continue
		}
		newTargets = append(newTargets, target)
	}
	fmt.Printf("chunk: %s\n", context2.GlobalChunkCtx.Address)
	for _, t := range newTargets{
		fmt.Printf("send to target: %s, dir: %s\n", t.DestAddr, t.DestDir)
	}

	if len(newTargets) == 0 {
		return nil
	}

	nextAddr := newTargets[0].DestAddr

	conn, _ := grpc.Dial(nextAddr, grpc.WithInsecure(), grpc.WithBlock())
	context2.GlobalChunkCtx.Client = pb.NewChunkServerClient(conn)

	stream, err := context2.GlobalChunkCtx.Client.PutObject(context.Background())

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	// Filter out itself
	info.Targets = newTargets
	pr := pb.PutRequest{
		Data: &pb.PutRequest_Info{
			Info: info,
		},
	}
	err = stream.Send(&pr)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	lastReadLen := 0
	fmt.Println(*writing)
	for {
		if (*buffer).Len() > lastReadLen {
			sendBuf := make([]byte, buffer.Len() - lastReadLen)
			(*buffer).Read(sendBuf)
			lastReadLen = (*buffer).Len()
			err = stream.Send(
				&pb.PutRequest{
					Data: &pb.PutRequest_Content{
						Content: sendBuf,
					},
				},
			)
			if err != nil {
				err = fmt.Errorf("failed to send chunkmanage via stream: %v", err)
				return err
			}

			if !*writing {
				break
			}
		}
	}

	_, e := stream.CloseAndRecv()

	fmt.Printf("stream %v\n", e)
	if e != nil {
		return e
	}

	return nil
}
