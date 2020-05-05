package object

import (
	"cherryfs/pkg/roles/dir"
	"cherryfs/pkg/roles/host"
)

type Object struct {
	Name string
	Size int64
	Hash string
	dir dir.Dir
	host host.Host
}
