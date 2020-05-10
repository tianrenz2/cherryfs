package chunk

import (
	"cherryfs/pkg/chunk/chunkserverpb"
	"google.golang.org/grpc"
	"bytes"
	"fmt"
	"context"
)

type ChunkContext struct {
	Address string
	Client  chunkserverpb.PutClient
}

func (chunkCtx *ChunkContext) SendAsRec(writing *bool, buffer *bytes.Buffer, info *chunkserverpb.ObjectInfo, targets []*chunkserverpb.Target) (error) {
	var newTargets = make([]*chunkserverpb.Target, 0)
	for _, target := range targets {
		if target.DestAddr == chunkCtx.Address {
			continue
		}
		newTargets = append(newTargets, target)
	}

	if len(newTargets) == 0 {
		return nil
	}

	nextAddr := newTargets[0].DestAddr

	conn, _ := grpc.Dial(nextAddr, grpc.WithInsecure(), grpc.WithBlock())
	chunkCtx.Client = chunkserverpb.NewPutClient(conn)

	stream, err := chunkCtx.Client.PutObject(context.Background())

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	// Filter out itself
	info.Targets = newTargets
	pr := chunkserverpb.PutRequest{
		Data: &chunkserverpb.PutRequest_Info{
			Info: info,
		},
	}
	fmt.Printf("Send to %s\n", nextAddr)
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
			//fmt.Printf("Read: %s, %d, %v", string(sendBuf), n, err)
			lastReadLen = (*buffer).Len()
			fmt.Printf("transfering: %v\n", string(sendBuf))
			err = stream.Send(
				&chunkserverpb.PutRequest{
					Data: &chunkserverpb.PutRequest_Content{
						Content: sendBuf,
					},
				},
			)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				err = fmt.Errorf("failed to send chunk via stream: %v", err)
				return err
			}

			if !*writing {
				break
			}
		}
	}

	ret, e := stream.CloseAndRecv()

	fmt.Printf("Response: %v\n", ret)
	if e != nil {
		return e
	}

	return nil
}
