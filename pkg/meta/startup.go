package meta

import (
	"cherryfs/pkg/meta/mgt"
	"cherryfs/pkg/meta/subgrouper"
	"fmt"
	rand2 "math/rand"
	"log"
)

var subgroupManager subgrouper.SubGroupManager = subgrouper.SubGroupManager{}

func Startup()  {
	clusterConfig := mgt.LoadConfig()
	hosts := mgt.LoadHosts(clusterConfig.Hosts)
	//subgroupManager.InitSubgroupSetup(hosts)
	subgrouper.GlobalSubGroupManager.InitSubgroupSetup(hosts)


	var spaceLevels = [5]int64{10, 20, 30, 40, 50}



	for _, subgroup := range subgrouper.GlobalSubGroupManager.SubGroups {
		for _, host := range subgroup.Hosts {
			for index, _ := range host.Dirs {
				sl := rand2.Int() % 5
				host.Dirs[index].TotalSpace = spaceLevels[sl]
				log.Printf("Dir %s has space %d\n", host.Dirs[index].Path, host.Dirs[index].TotalSpace)
			}
		}

	}



	fmt.Println(subgroupManager.SubGroups)
	//server.StartServer()
}
