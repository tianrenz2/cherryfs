package mgt

import (
	"cherryfs/pkg/meta/subgroup"
)

type GlobalConfig struct {
	SubGroupNum int
	SubGroups []subgroup.SubGroup
}

