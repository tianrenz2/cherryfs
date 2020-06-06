package context

import "cherryfs/pkg/roles/dir"

type ChunkInfo struct {
	HostId string
	Addr string
	Dirs []dir.Dir
}
