package mock

import (
	"cherryfs/pkg/context"
	"cherryfs/pkg/meta/initialize"
	"cherryfs/pkg/meta/subgroup"
	"cherryfs/pkg/role/dir"
	"cherryfs/pkg/role/host"
)

func MockGlobalCtx() *context.Context {
	globalCtx := initialize.Startup()
	return globalCtx
}

func MockSgManager() *subgroup.SubGroupManager {

	return nil
}

func MockHostManager() *host.HostManager {

	return nil
}

func MockDirManager() *dir.DirManager {

	return nil
}
