package host

import (
	"fmt"
	"github.com/google/uuid"
	"cherryfs/pkg/roles/dir"
)

type ConfigHost struct {
	Hostname string
	Address string
	Dirs []string
}

type Host struct {
	SubgroupId int
	HostId string
	Hostname string
	Address string
	Dirs []string
}

/*
	HostManager is responsible for managing all hosts in the cluster
*/
type HostManager struct {
	Hosts []*Host
	hostMap map[string]*Host
}

func (hostMg *HostManager) PrintHostMap() {
	fmt.Println(hostMg.hostMap)
}

func (hostMg *HostManager) GetHostByHostId(HostId string) (Host, error) {
	if host, ok := hostMg.hostMap[HostId]; ok {
		return *host, nil
	}
	return Host{}, fmt.Errorf("could not find host for %s", HostId)
}

func (hostMg *HostManager) InitHostMap() (error) {
	hostMg.hostMap = make(map[string]*Host)
	for _, host := range hostMg.Hosts {
		hostMg.hostMap[host.HostId] = host
	}

	return nil
}

func (hostMg *HostManager) InitAllHosts(configHosts []ConfigHost, dirManager *dir.DirManager) (error) {
	hostMg.Hosts = make([]*Host, 0)
	for _, configHost := range configHosts {
		dirIds := make([]string, 0)
		hostId := uuid.New().String()
		for _, d := range configHost.Dirs {
			id, _ := dirManager.CreateDir(d, hostId)
			dirIds = append(dirIds, id)
		}

		hostMg.Hosts = append(hostMg.Hosts, &Host{
			HostId:hostId,
			Hostname: configHost.Hostname,
			Address: configHost.Address,
			Dirs:dirIds,
		})
	}
	hostMg.InitHostMap()
	dirManager.InitDirMap()

	return nil
}

func (hostMg *HostManager) AddHost(hostAddr string, hostDirs []dir.Dir, dirManager *dir.DirManager) (string, error) {
	host, err := hostMg.InitNewHost(hostAddr, hostDirs, dirManager)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	hostMg.Hosts = append(hostMg.Hosts, host)
	hostMg.hostMap[host.HostId] = host
	return host.HostId, nil
}

func (hostMg *HostManager) InitNewHost(hostAddr string, hostDirs []dir.Dir, dirManager *dir.DirManager) (*Host, error)  {
	hostId := uuid.New().String()
	dirs := make([]string, 0)
	for _, hostDir := range hostDirs {
		id, _ := dirManager.CreateDir(hostDir.Path, hostId)
		dirManager.SetTotalSpace(id, hostDir.TotalSpace)
		dirManager.SetUsedSpace(id, hostDir.UsedSpace)
		dirs = append(dirs, id)
	}

	host := Host{
		HostId:hostId,
		Address:hostAddr,
		Dirs:dirs,
	}

	return &host, nil
}