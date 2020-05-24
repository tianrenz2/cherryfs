package comm

import (
	"cherryfs/pkg/object"
	"cherryfs/pkg/comm/chunkserverpb"
)

type ObjectSender struct {
		
}

func (objectSender *ObjectSender) SendObject(object object.LocalObject, info chunkserverpb.ObjectInfo, target chunkserverpb.Target) (error) {


	return nil
}

