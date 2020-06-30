package mock

import (
	"cherryfs/pkg/context"
	"cherryfs/pkg/meta/subgroup"
	"cherryfs/pkg/roles/host"
	"cherryfs/pkg/roles/dir"
	"cherryfs/pkg/meta/initialize"
)

func MockGlobalCtx() *context.Context {
	globalCtx := initialize.Startup()
	return globalCtx
}


func MockSgManager() *subgroup.SubGroupManager {

	return nil
}


func MockHostManager() *host.HostManager{

	return nil
}


func MockDirManager() *dir.DirManager {

	return nil
}