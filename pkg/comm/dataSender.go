package comm

import (
	"cherryfs/pkg/object"
	"cherryfs/pkg/chunk/chunkserverpb"
)

type ObjectSender struct {
		
}

func (objectSender *ObjectSender) SendObject(object object.LocalObject, info chunkserverpb.ObjectInfo, target chunkserverpb.Target) (error) {


	return nil
}

