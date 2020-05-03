package allocator

import (
	"cherryfs/pkg/meta/subgrouper"
	"crypto/sha256"
	"encoding/binary"
)

type RedundancyPolicy int

const (
	ReplicaPolicy RedundancyPolicy = 1
	ReplicNum int = 3
)

type Allocator struct {
	policy RedundancyPolicy

}

func (allocator *Allocator) AllocSubgroups(objName, objHash string, objSize int) ([]subgrouper.SubGroup, error) {
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
	}

	return allocatedSubGroups, nil
}
