package allocator

import (
	"cherryfs/pkg/meta/subgroup"
	"crypto/sha256"
	"encoding/binary"
	"cherryfs/pkg/object"
)

/*
	Responsible for allocating subgroups to an object, then from the each
	selected subgroup, a target will be selected to place the object.
*/

type RedundancyPolicy int

const (
	ReplicaPolicy RedundancyPolicy = 1
	ReplicaNum    int              = 3
)

type Allocator struct {
	Policy RedundancyPolicy

}

func (allocator *Allocator) AllocSubgroups(object object.Object) ([]subgroup.SubGroup, error) {
	objName := object.Name
	keyByte := []byte(objName)
	h := sha256.New()
	h.Write(keyByte)
	bs := h.Sum(nil)
	hashNum := binary.BigEndian.Uint64(bs)

	subGroupNum := subgroup.GlobalSubGroupManager.GetSubGroupNumber()

	modStart := int(hashNum) % subGroupNum

	var allocatedSubGroups = make([]subgroup.SubGroup, 0)

	for i := 0; i < ReplicaNum; i ++ {
		allocatedSubGroups = append(allocatedSubGroups, subgroup.GlobalSubGroupManager.GetSubGroupById(modStart))
		modStart += 1
		modStart %= subGroupNum
	}

	return allocatedSubGroups, nil
}
