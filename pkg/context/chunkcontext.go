package context

import "cherryfs/pkg/roles/dir"

type ChunkInfo struct {
	Addr string
	Dirs []dir.Dir
}
