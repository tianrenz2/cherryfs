package host

import "cherryfs/pkg/roles/dir"

type Host struct {
	Hostname string
	Address string
	Dirs []dir.Dir
}

