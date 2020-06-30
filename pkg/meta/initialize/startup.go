package initialize

import (
	"cherryfs/pkg/roles/dir"
	"cherryfs/pkg/roles/host"
	"cherryfs/pkg/meta/subgroup"
	"fmt"
	"cherryfs/pkg/context"
	"math/rand"
	"os"
	"cherryfs/pkg/etcd"
)

func Startup() (*context.Context) {

	var dirSpaces = []int64 {1, 5, 7, 10}

	clusterConfig := LoadConfig()
	var dirManager dir.DirManager
	var hostManager host.HostManager
	var GlobalSubGroupManager subgroup.SubGroupManager

	dirManager.New()

	hostManager.InitAllHosts(clusterConfig.Hosts, &dirManager)
	ctx := context.Context{SGManager: &GlobalSubGroupManager, HManager:&hostManager, DManager:&dirManager}

	ctx.SGManager.InitSubgroupSetup(hostManager.Hosts)

	for _, sg := range ctx.SGManager.SubGroups {
		fmt.Printf("sg: %d\n", sg.SubGroupId)
		for _, hId := range sg.Hosts {
			fmt.Printf("host: %s\n", hId)
			h, _ := hostManager.GetHostByHostId(hId)
			for _, d := range h.Dirs {
				fmt.Printf("dir: %s\n", d)
				//dname, _ := (&dirManager).GetDirByDirId(d)
				dirManager.UpdateDirSpaceByDirId(d, dirSpaces[rand.Int() % 4] * 1e6)
				//fmt.Printf("updated space: %d, %d\n", d2.TotalSpace, dname.TotalSpace)
			}
		}
	}

	ctx.EtcdCli.CreateEtcdClient(os.Getenv("ETCDADDR"))
	//err := ctx.PersistCluster()

	//if err != nil {
	//	fmt.Errorf("%v", err)
	//}
	fmt.Printf("created cluster successfully")

	return &ctx
}

func LoadClusterConfig() (context.Context) {
	ctx := context.Context{}
	ctx.InitManagers()
	ctx.EtcdCli = etcd.EtcdClient{}
	ctx.EtcdCli.CreateEtcdClient(os.Getenv("ETCDADDR"))

	//ctx.EtcdCli.CreateEtcdClient(os.Getenv("ETCDADDR"))
	err := ctx.RecoverCluster()

	if err != nil{
		fmt.Errorf("%v", err)
	}

	return ctx
}

func main()  {
	Startup()
	LoadClusterConfig()
}