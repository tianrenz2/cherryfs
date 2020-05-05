package meta

import (
	"cherryfs/pkg/meta/mgt"
	"cherryfs/pkg/roles/dir"
	"cherryfs/pkg/roles/host"
	"cherryfs/pkg/meta/subgroup"
	"fmt"
	"cherryfs/pkg/context"
	"math/rand"
)

func Startup() (context.Context) {

	var dirSpaces = []int64 {1, 5, 7, 10}

	clusterConfig := mgt.LoadConfig()
	var dirManager dir.DirManager
	var hostManager host.HostManager
	var GlobalSubGroupManager subgroup.SubGroupManager

	hostManager.InitAllHosts(clusterConfig.Hosts, &dirManager)
	ctx := context.Context{SGManager: &GlobalSubGroupManager, HManager:&hostManager, DManager:&dirManager}

	ctx.SGManager.InitSubgroupSetup(hostManager.Hosts)

	for _, sg := range ctx.SGManager.SubGroups {
		for _, hId := range sg.Hosts {
			h, _ := hostManager.GetHostByHostId(hId)
			for _, d := range h.Dirs {
				//dname, _ := (&dirManager).GetDirByDirId(d)
				dirManager.UpdateDirSpaceByDirId(d, dirSpaces[rand.Int() % 4] * 1e6)
				//fmt.Printf("updated space: %d, %d\n", d2.TotalSpace, dname.TotalSpace)
			}
		}
	}


	//fmt.Println(ctx.HManager.Hosts)
	ctx.EtcdCli.CreateEtcdClient("127.0.0.1", 2380)
	err := ctx.PersistCluster()

	if err != nil {
		fmt.Errorf("%v", err)
	}
	fmt.Printf("created cluster successfully")

	return ctx
}

func LoadClusterConfig() (context.Context) {
	ctx := context.Context{}
	ctx.InitManagers()
	ctx.EtcdCli.CreateEtcdClient("127.0.0.1", 2380)
	err := ctx.RecoverCluster()

	if err != nil{
		fmt.Errorf("%v", err)
	}

	return ctx
}

func main()  {
	Startup()
	//LoadClusterConfig()
}