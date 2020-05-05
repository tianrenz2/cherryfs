package allocator

import (
	"cherryfs/pkg/meta/subgroup"
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

func (allocator *Allocator) AllocateTargetsFromSgs(subgroups []subgroup.SubGroup, obj object.Object) ([]Target, error) {
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

func (allocator *Allocator) AllocateTargetFromSg(subgroup subgroup.SubGroup, obj object.Object) (Target, error) {
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