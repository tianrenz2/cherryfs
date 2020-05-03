package subgrouper

import (
	cfg "cherryfs/pkg/config"
	"cherryfs/pkg/roles/host"
	"cherryfs/pkg/roles/dir"
	"fmt"
)

type SubGroup struct {
	SubGroupId int
	Hosts []host.Host
	DManager dir.DirManager
}

type SubGroupManager struct {
	SubGroups []SubGroup
}

var GlobalSubGroupManager SubGroupManager

func (subGroupMg *SubGroupManager) InitSubgroupSetup(allHosts []host.Host) (error) {

	subgroups, err := subGroupMg.InitSubgroups(allHosts)
	if err != nil{
		return fmt.Errorf("failed to init subgroups: %v", err)
	}
	subGroupMg.SubGroups = subgroups
	return nil
}

func (subGroupMg *SubGroupManager) GetSubGroupNumber() int {
	return len(subGroupMg.SubGroups)
}

func (subGroupMg *SubGroupManager) GetSubGroupById(SubgroupId int) SubGroup {
	return subGroupMg.SubGroups[SubgroupId]
}

func (subGroupMg *SubGroupManager)InitSubgroups(allHosts []host.Host) ([]SubGroup, error){
	var subgroupNum = 0
	hostNum := len(allHosts)

	if cfg.MinReplicaNum > hostNum {
		subgroupNum = len(allHosts)
	} else {
		subgroupNum = cfg.MinReplicaNum
	}

	var subgroups = make([]SubGroup, 0)
	numPerGroup := hostNum / subgroupNum + 1

	for groupIndex :=0; groupIndex < subgroupNum; groupIndex++ {
		groupStart := groupIndex * numPerGroup
		groupEnd := groupStart + numPerGroup
		subgroup, err := subGroupMg.InitOneSubgroup(allHosts, groupIndex, groupStart, groupEnd)
		if err != nil {
			return subgroups, fmt.Errorf("failed to init subgroups: %v", err)
		}

		subgroups = append(subgroups, subgroup)
	}

	return subgroups, nil
}

func (subGroupMg *SubGroupManager) InitOneSubgroup(allHosts []host.Host, groupIndex, groupStart, groupEnd int) (SubGroup, error) {
	var subgroup = SubGroup{Hosts: make([]host.Host, 0), SubGroupId:groupIndex}
	subgroup.DManager = dir.DirManager{ReliefNum: dir.DefaultReliefNum}

	if groupEnd > len(allHosts) {
		groupEnd = len(allHosts)
	}
	for hostIndex:= groupStart; hostIndex < groupEnd; hostIndex++ {
		for i, _ := range allHosts[hostIndex].Dirs {
			allHosts[hostIndex].Dirs[i].Manager = subgroup.DManager
		}
		subgroup.Hosts = append(subgroup.Hosts, allHosts[hostIndex])
	}

	return subgroup, nil
}