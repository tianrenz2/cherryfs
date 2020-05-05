package allocator

import (
	"cherryfs/pkg/meta/subgroup"
	"crypto/sha256"
	"encoding/binary"
	"cherryfs/pkg/object"
	"cherryfs/pkg/context"
	"fmt"
)

/*
	SubGroup Allocator: Responsible for allocating subgroups to an object, then
	from the each selected subgroup, a target will be selected to place the object.
*/

type RedundancyPolicy int

const (
	ReplicaPolicy RedundancyPolicy = 1
	ReplicaNum    int              = 3
)

type Allocator struct {
	Policy RedundancyPolicy
	Ctx context.Context
}

func (allocator *Allocator) AllocSubgroups(object object.Object) ([]subgroup.SubGroup, error) {
	objName := object.Name
	keyByte := []byte(objName)
	h := sha256.New()
	h.Write(keyByte)
	bs := h.Sum(nil)
	hashNum := binary.BigEndian.Uint64(bs)

	subGroupNum := allocator.Ctx.SGManager.GetSubGroupNumber()

	if subGroupNum <= 0 {
		return make([]subgroup.SubGroup, 0), fmt.Errorf("there must be at least 1 subgroup")
	}

	modStart := int(hashNum) % subGroupNum

	var allocatedSubGroups = make([]subgroup.SubGroup, 0)

	for i := 0; i < ReplicaNum; i ++ {
		allocatedSubGroups = append(allocatedSubGroups, allocator.Ctx.SGManager.GetSubGroupById(modStart))
		modStart += 1
		modStart %= subGroupNum
	}

	return allocatedSubGroups, nil
}
