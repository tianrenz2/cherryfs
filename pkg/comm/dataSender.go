package comm

import (
	"cherryfs/pkg/object"
	"cherryfs/pkg/comm/pb"
)

type ObjectSender struct {
		
}

func (objectSender *ObjectSender) SendObject(object object.LocalObject, info chunkserverpb.ObjectInfo, target chunkserverpb.Target) (error) {


	return nil
}

