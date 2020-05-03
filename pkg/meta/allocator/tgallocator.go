package allocator

import (
	"cherryfs/pkg/meta/subgrouper"
	"cherryfs/pkg/roles/dir"
	"cherryfs/pkg/roles/host"
	"cherryfs/pkg/object"
	"fmt"
)

/*
	Target Allocator: this part is responsible of allocating targets for
	a client's put request, each target contains a Dir & Host.
*/

type Target struct {
	Host host.Host
	Dir dir.Dir
}

func (allocator *Allocator) AllocateTargetsFromSgs(subgroups []subgrouper.SubGroup, obj object.Object) ([]Target, error) {
	var allocatedTgs = make([]Target, 0)
	for _, subgroup := range subgroups {
		target, err := allocator.AllocateTargetFromSg(subgroup, obj)
		if err != nil{
			return allocatedTgs, err
		}
		allocatedTgs = append(allocatedTgs, target)
	}
	return allocatedTgs, nil
}

func (allocator *Allocator) AllocateTargetFromSg(subgroup subgrouper.SubGroup, obj object.Object) (Target, error) {
	var selectedHost interface{} = nil
	var selectedDir interface{} = nil
	var maxScore float64 = 0

	target := Target{}

	for _, h := range subgroup.Hosts {
		for _, d := range h.Dirs {
			baseScore := d.GetBaseScore()
			multiplier := d.TotalSpace - d.UsedSpace - obj.Size
			score := baseScore * float64(multiplier)
			if score > maxScore {
				maxScore = score
				selectedHost = h
				selectedDir = d
			}
		}
	}
	if selectedHost == nil || selectedDir == nil {
		return target, fmt.Errorf("Did not find qualified dir in subgroup %d", subgroup.SubGroupId)
	}

	target.Host = selectedHost.(host.Host)
	target.Dir = selectedDir.(dir.Dir)
	return target, nil
}