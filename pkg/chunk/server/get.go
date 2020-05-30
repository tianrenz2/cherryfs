package main

import (
	"cherryfs/pkg/comm/pb"
	"os"
	"io"
	"fmt"
)

func (s *ChunkServer) GetObject(getRequest *pb.GetRequest, responser pb.ChunkServer_GetObjectServer) (error) {

	dir := getRequest.Dir
	name := getRequest.Name

	f, _ := os.Open(dir + "/" + name)
	buf := make([]byte, 1024)

	sending := true

	for sending {
		n, err := f.Read(buf)

		if err != nil {
			if err == io.EOF {
				sending = false
				err = nil
				continue
			}
			err = fmt.Errorf("errored while copying from file to buf")
		}


		err = responser.Send(
			& pb.GetResponse{
				Content: buf[:n],
			},
		)

		if err != nil {
			err = fmt.Errorf("failed to send chunkmanage via stream: %v", err)
			return err
		}
	}

	return nil
}