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
	for _, sg := range subgroups {
		target, err := allocator.AllocateTargetFromSg(sg, obj)
		if err != nil{
			return allocatedTgs, err
		}
		allocatedTgs = append(allocatedTgs, target)
	}
	return allocatedTgs, nil
}

func (allocator *Allocator) AllocateTargetFromSg(subgroup subgroup.SubGroup, obj object.Object) (Target, error) {
	var selectedHostId interface{} = nil
	var selectedDirId interface{} = nil
	var maxScore float64 = 0

	target := Target{}

	for _, hId := range subgroup.Hosts {
		h, err := allocator.Ctx.HManager.GetHostByHostId(hId)
		if err != nil {
			return Target{}, fmt.Errorf("%v", err)
		}

		for _, dId := range h.Dirs {
			var d, err = allocator.Ctx.DManager.GetDirByDirId(dId)
			if err != nil {
				return Target{}, fmt.Errorf("%v", err)
			}

			baseScore := d.GetBaseScore()
			multiplier := d.TotalSpace - d.UsedSpace - obj.Size
			score := baseScore * float64(multiplier)
			if score > maxScore {
				maxScore = score
				selectedHostId = hId
				selectedDirId = dId
			}
		}
	}
	if selectedHostId == nil || selectedDirId == nil {
		return target, fmt.Errorf("did not find qualified dir in subgroup %d", subgroup.SubGroupId)
	}

	selectedHost, err := allocator.Ctx.HManager.GetHostByHostId(fmt.Sprintf("%v", selectedHostId))

	selectedDir, err := allocator.Ctx.DManager.GetDirByDirId(fmt.Sprintf("%v", selectedDirId))

	if err != nil {
		return Target{}, fmt.Errorf("%v", err)
	}
	target.Host = selectedHost
	target.Dir = selectedDir

	return target, nil
}