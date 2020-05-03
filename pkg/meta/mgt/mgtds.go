package mgt

import (
	"cherryfs/pkg/meta/subgrouper"
)

type GlobalConfig struct {
	SubGroupNum int
	SubGroups []subgrouper.SubGroup
}

