package mgt

import (
	"cherryfs/pkg/meta/subgrouper"
	"cherryfs/pkg/roles/dir"
)

type Host struct {
	Hostname string
	Address string
	Dirs []dir.Dir
}

type GlobalConfig struct {
	SubGroupNum int
	SubGroups []subgrouper.SubGroup
}

