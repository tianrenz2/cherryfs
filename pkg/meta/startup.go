package main

import (
	"cherryfs/pkg/meta/mgt"
	"cherryfs/pkg/roles/dir"
	"cherryfs/pkg/roles/host"
	"cherryfs/pkg/meta/subgroup"
	"fmt"
	"cherryfs/pkg/meta/context"
)

func Startup()  {
	clusterConfig := mgt.LoadConfig()
	var dirManager dir.DirManager
	var hostManager host.HostManager

	hostManager.InitAllHosts(clusterConfig.Hosts, &dirManager)

	subgroup.GlobalSubGroupManager.InitSubgroupSetup(hostManager.Hosts)

	for _, sg := range subgroup.GlobalSubGroupManager.SubGroups {
		for _, hId := range sg.Hosts {
			h, _ := hostManager.GetHostByHostId(hId)
			for _, d := range h.Dirs {
				dname, _ := dirManager.GetDirByDirId(d)
				fmt.Println(dname)
			}
		}
	}

	ctx := context.Context{SGManager: subgroup.GlobalSubGroupManager, HManager:hostManager, DManager:dirManager}
	ctx.EtcdCli.CreateEtcdClient("127.0.0.1", 2380)
	err := ctx.PersistClusterConfig()

	if err != nil {
		fmt.Errorf("%v", err)
	}
	fmt.Printf("created cluster successfully")
	//fmt.Println(subgroup.GlobalSubGroupManager.SubGroups)
	//m, _ := json.Marshal(subgroup.GlobalSubGroupManager.SubGroups)
	//g := base64.StdEncoding.EncodeToString(m)

	//fmt.Println(g)
	//var b []byte
	//b, _ = base64.StdEncoding.DecodeString(g)
	//
	//var s []subgroup.SubGroup
	//json.Unmarshal(b, &s)
	//fmt.Println(s)

	//server.StartServer()
}

func LoadClusterConfig() {
	ctx := context.Context{}
	ctx.EtcdCli.CreateEtcdClient("127.0.0.1", 2380)
	err := ctx.DecodeSubgroups()

	if err != nil{
		fmt.Errorf("%v", err)
	}

	fmt.Println(ctx.SGManager.SubGroups)
}

func main()  {
	//Startup()
	LoadClusterConfig()
}