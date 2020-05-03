package subgrouper

import (
	"cherryfs/pkg/meta/mgt"
	cfg "cherryfs/pkg/config"
)

type SubGroup struct {
	SubGroupId int
	Hosts []mgt.Host
}

type SubGroupManager struct {
	SubGroups []SubGroup
}

var GlobalSubGroupManager SubGroupManager

func (subGroupMg *SubGroupManager) InitSubgroupSetup(allHosts []mgt.Host) {
	subGroupMg.SubGroups = InitSubgroups(allHosts)
}

func (subGroupMg *SubGroupManager) GetSubGroupNumber() int {
	return len(subGroupMg.SubGroups)
}

func (subGroupMg *SubGroupManager) GetSubGroupById(SubgroupId int) SubGroup {
	return subGroupMg.SubGroups[SubgroupId]
}

func InitSubgroups(allHosts []mgt.Host) []SubGroup {
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
		var subgroup = SubGroup{Hosts: make([]mgt.Host, 0), SubGroupId:groupIndex}
		splitStart := groupIndex * numPerGroup
		splitEnd := splitStart + numPerGroup

		if splitEnd > len(allHosts) {
			splitEnd = len(allHosts)
		}
		for hostIndex:=splitStart; hostIndex < splitEnd; hostIndex++ {
			subgroup.Hosts = append(subgroup.Hosts, allHosts[hostIndex])
		}
		subgroups = append(subgroups, subgroup)
	}

	return subgroups
}
