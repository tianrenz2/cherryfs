package watchservice_test

import (
	"testing"
	"cherryfs/pkg/mock"
	"github.com/stretchr/testify/assert"
	"fmt"
	"cherryfs/pkg/object"
	"cherryfs/pkg/comm/pb"
	"strconv"
	"cherryfs/pkg/meta/watchservice"
)


func TestGenerateRecoverMap(t *testing.T) {

	globalCtx := mock.MockGlobalCtx()

	assert.Equal(t, globalCtx.SGManager.GetSubGroupNumber(), 3, "subgroup number should be 3")

	fmt.Println(globalCtx.HManager.Hosts)
	var mockObjects = make([]object.Object, 0)
	//
	targets := make([]*pb.Target,0)

	for i:=0; i < 3; i++ {
		targetId := globalCtx.SGManager.SubGroups[i].Hosts[0]
		tar, _ := globalCtx.HManager.GetHostByHostId(targetId)

		targets = append(targets, &pb.Target{
			DestAddr: tar.Address,
			DestId: tar.HostId,
			DestDir: tar.Dirs[0],
			SgId: int32(i),
		})
	}

	objNumber := 4

	for i:=0; i < objNumber; i++ {
		mockObjects = append(mockObjects, object.Object{
			Name: "obj-" + strconv.Itoa(i),
			Size: 1,
			Hash: "vvv",
			Targets: targets,
		})
	}


	recoverMap, err := watchservice.GenerateRecoverMap(globalCtx, 0, 1, mockObjects)
	assert.Equal(t, err, nil, "should not yield error")

	recoverObjNum := 0

	for _, copyInfoList := range recoverMap{
		for _, copyInfo := range copyInfoList {
			recoverObjNum += len(copyInfo.LocalObjects)
		}
	}

	assert.Equal(t, recoverObjNum, objNumber, "not all lost objects are set to be recovered")

	return
}