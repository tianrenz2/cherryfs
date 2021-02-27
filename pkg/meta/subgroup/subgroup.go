package subgroup

import (
	"cherryfs/pkg/config"
	"cherryfs/pkg/role/dir"
	"cherryfs/pkg/role/host"
	"fmt"
)

/*
	A Subgroup (or referenced as 'sg' in some other places) is a unit not directly
	visible to users. A subgroup contains multiple hosts, usually 1 subgroup has
	1 replica of an object.
*/

type SubGroup struct {
	SubGroupId  int
	Hosts       []string
	DirSettings dir.DirSubGroupSetting
}

type SubGroupManager struct {
	SubGroups []*SubGroup
}

func (subGroupMg *SubGroupManager) InitSubgroupSetup(allHosts []*host.Host) error {

	subgroups, err := subGroupMg.InitSubgroups(allHosts)
	if err != nil {
		return fmt.Errorf("failed to initialize subgroups: %v", err)
	}
	subGroupMg.SubGroups = subgroups
	return nil
}

func (subGroupMg *SubGroupManager) GetSubGroupNumber() int {
	return len(subGroupMg.SubGroups)
}

func (subGroupMg *SubGroupManager) GetSubGroupById(SubgroupId int) SubGroup {
	return *subGroupMg.SubGroups[SubgroupId]
}

func (subGroupMg *SubGroupManager) InitSubgroups(allHosts []*host.Host) ([]*SubGroup, error) {
	var subgroupNum = 0
	hostNum := len(allHosts)

	if hostNum <= config.MinReplicaNum {
		subgroupNum = len(allHosts)
	} else {
		subgroupNum = config.MinReplicaNum
	}

	var subgroups = make([]*SubGroup, 0)

	if subgroupNum == 0 {
		return subgroups, nil
	}

	numPerGroup := hostNum / subgroupNum

	for groupIndex := 0; groupIndex < subgroupNum; groupIndex++ {
		groupStart := groupIndex * numPerGroup
		groupEnd := groupStart + numPerGroup
		subgroup, err := subGroupMg.InitOneSubgroup(allHosts, groupIndex, groupStart, groupEnd)
		if err != nil {
			return subgroups, fmt.Errorf("failed to initialize subgroups: %v", err)
		}

		subgroups = append(subgroups, &subgroup)
	}

	return subgroups, nil
}

func (subGroupMg *SubGroupManager) InitOneSubgroup(allHosts []*host.Host, groupIndex, groupStart, groupEnd int) (SubGroup, error) {
	var subgroup = SubGroup{Hosts: make([]string, 0), SubGroupId: groupIndex}
	subgroup.DirSettings = dir.DirSubGroupSetting{ReliefNum: dir.DefaultReliefNum}

	if groupEnd > len(allHosts) {
		groupEnd = len(allHosts)
	}

	for hostIndex := groupStart; hostIndex < groupEnd; hostIndex++ {
		subgroup.Hosts = append(subgroup.Hosts, allHosts[hostIndex].HostId)
	}

	return subgroup, nil
}

func (subGroupMg *SubGroupManager) AssignSubgroupToHost() {
	if len(subGroupMg.SubGroups) < config.MinReplicaNum {

	}
}

func (subGroupMg *SubGroupManager) AddSubgoup() {
	groupIndex := len(subGroupMg.SubGroups)
	var subgroup = SubGroup{Hosts: make([]string, 0), SubGroupId: groupIndex}
	subGroupMg.SubGroups = append(subGroupMg.SubGroups, &subgroup)
}

func (subgroup *SubGroup) AddHost(host host.Host) {
	subgroup.Hosts = append(subgroup.Hosts, host.HostId)
}
