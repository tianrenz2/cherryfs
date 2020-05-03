package allocator

import (
	"cherryfs/pkg/meta/subgrouper"
	"crypto/sha256"
	"encoding/binary"
	"cherryfs/pkg/object"
)

type RedundancyPolicy int

const (
	ReplicaPolicy RedundancyPolicy = 1
	ReplicNum int = 3
)

type Allocator struct {
	Policy RedundancyPolicy

}

func (allocator *Allocator) AllocTargets(object object.Object) ([]Target, error) {
	allocSgs, errSg := allocator.AllocSubgroups(object)
	var targets = make([]Target,0)

	if errSg != nil {
		return targets, errSg
	}
	targets, errTg := allocator.AllocateTargetsFromSgs(allocSgs, object)
	if errTg != nil {
		return targets, errTg
	}

	return targets, nil
}

func (allocator *Allocator) AllocSubgroups(object object.Object) ([]subgrouper.SubGroup, error) {
	objName := object.Name
	keyByte := []byte(objName)
	h := sha256.New()
	h.Write(keyByte)
	bs := h.Sum(nil)
	hashNum := binary.BigEndian.Uint64(bs)

	subGroupNum := subgrouper.GlobalSubGroupManager.GetSubGroupNumber()

	modStart := int(hashNum) % subGroupNum

	var allocatedSubGroups = make([]subgrouper.SubGroup, 0)

	for i := 0; i < ReplicNum; i ++ {
		allocatedSubGroups = append(allocatedSubGroups, subgrouper.GlobalSubGroupManager.GetSubGroupById(modStart))
		modStart += 1
		modStart %= subGroupNum
	}

	return allocatedSubGroups, nil
}
