package allocator

import (
	"cherryfs/pkg/meta/subgrouper"
	"cherryfs/pkg/meta/mgt"
)

/*
	Target Allocator: this part is responsible of allocating targets for
	a client's put request, each target contains a Dir & Host.
*/

type Target struct {
	host mgt.Host
	dir mgt.Dir
}


func (allocator *Allocator) AllocateTargets(subgroups []subgrouper.SubGroup) ([]Target, error) {
	var allocatedTgs = make([]Target, 0)
	for _, subgroup := range subgroups {
		target, err := allocator.AllocateTargetFromSg(subgroup)
		if err != nil{
			return allocatedTgs, err
		}
		allocatedTgs = append(allocatedTgs, target)
	}
	return allocatedTgs, nil
}


func (allocator *Allocator) AllocateTargetFromSg(subgroup subgrouper.SubGroup) (Target, error) {
	var selectedHost = nil
	var selectedDir = nil
	var maxScore = 0

	for _, host := range subgroup.Hosts {
		for _, dir := range host.Dirs {

		}
	}

}