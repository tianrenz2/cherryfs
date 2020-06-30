package watchservice

import (
	"cherryfs/pkg/context"
	"cherryfs/pkg/object"
	"cherryfs/pkg/task"
	"cherryfs/pkg/comm/pb"
	"cherryfs/pkg/meta/allocator"
	"cherryfs/pkg/client"
	"log"
	"encoding/json"
	"fmt"
)


func GenerateRecoverMap(ctx *context.Context, srcSgId, destSgId int, objects []object.Object) (map[string][]*task.CopyTaskInfo, error) {
	sg := ctx.SGManager.GetSubGroupById(destSgId)
	targetObjectMap := make(map[string][]object.Object)

	fmt.Printf("dest sg: %d \n", sg.SubGroupId)

	for _, obj := range objects {
		//log.Printf("objects to recover: %v", obj.Name)
		reAllocator := allocator.Allocator{
			Policy: allocator.ReplicaPolicy,
			Ctx: *ctx,
		}
		newTarget, err := reAllocator.AllocateTargetFromSg(sg, obj)

		log.Printf("new target for object %s is sg %d %s\n", obj.Name, newTarget.SgId, newTarget.Host.HostId)

		if err != nil {
			log.Printf("failed to assign a new target for object %s, %v", obj.Name, err)
		}

		newTargetByte, _ := json.Marshal(newTarget)
		newTargetStr := string(newTargetByte)

		if _, ok := targetObjectMap[newTargetStr]; !ok {
			targetObjectMap[newTargetStr] = make([]object.Object, 0)
		}
		targetObjectMap[newTargetStr] = append(targetObjectMap[newTargetStr], obj)
	}

	type TarObjPair struct {
		HostId string
		Objects []object.Object
	}

	addrInfoMap := make(map[string][]*task.CopyTaskInfo)
	for newTargetStr, objects := range targetObjectMap {
		var newTarget allocator.Target
		json.Unmarshal([]byte(newTargetStr), &newTarget)

		for _, obj := range objects {
			log.Printf("obj targets: %v \n", obj.Targets)
			var selectedSrcTarget pb.Target
			for _, target := range obj.Targets{
				if int(target.SgId) == srcSgId {
					selectedSrcTarget = *target
					break
				}
			}
			if _, ok := addrInfoMap[selectedSrcTarget.DestAddr]; !ok {
				addrInfoMap[selectedSrcTarget.DestAddr] = make([]*task.CopyTaskInfo, 0)
				addrInfoMap[selectedSrcTarget.DestAddr] = append(
					addrInfoMap[selectedSrcTarget.DestAddr],
						&task.CopyTaskInfo{
							Target: pb.Target{
								DestId: newTarget.Host.HostId,
								DestDir: newTarget.Dir.Path,
								DestAddr: newTarget.Host.Address,
								SgId: int32(newTarget.SgId),
							},
					},
				)
			} else {
				existingCopyTaskInfoLen := len(addrInfoMap[selectedSrcTarget.DestAddr])
				if existingCopyTaskInfoLen > 0 && addrInfoMap[selectedSrcTarget.DestAddr][existingCopyTaskInfoLen - 1].Target.DestId != newTarget.Host.HostId{
					addrInfoMap[selectedSrcTarget.DestAddr] = append(
						addrInfoMap[selectedSrcTarget.DestAddr],
						&task.CopyTaskInfo{
							Target: pb.Target{
								DestId: newTarget.Host.HostId,
								DestDir: newTarget.Dir.Path,
								DestAddr: newTarget.Host.Address,
								SgId: int32(newTarget.SgId),
							},
						},
					)
				}
			}

			addrInfoMap[selectedSrcTarget.DestAddr][len(addrInfoMap[selectedSrcTarget.DestAddr]) - 1].AddLcObject(
				object.LocalObject{
					Name: obj.Name,
					Size: obj.Size,
					Hash: obj.Hash,
					Path: selectedSrcTarget.DestDir,
					HostId: selectedSrcTarget.DestId,
				},
			)
		}
	}

	return addrInfoMap, nil
}


func SetupRecoverTransmissions(recoverMap map[string][]*task.CopyTaskInfo) error {
	for senderAddr, copyInfoPtrList := range recoverMap {

		sendCopyInfoList := make([]task.CopyTaskInfo, 0)
		for _, copyInfoPtr := range copyInfoPtrList {
			sendCopyInfoList = append(sendCopyInfoList, *copyInfoPtr)
		}
		log.Printf("sending from %s, number of targets %d\n", senderAddr, len(sendCopyInfoList))
		senderClient := client.InternalClient{}
		senderClient.New(senderAddr)
		err := senderClient.SendTask(task.CopyObjects, sendCopyInfoList)
		if err != nil {
			log.Printf("failed to set up transmission task for chunk %s\n", senderAddr)
		}
		log.Printf("successfully set up task for recovering from host %s \n", senderAddr)
	}

	return nil
}
