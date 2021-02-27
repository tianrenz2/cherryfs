package host

import (
	"cherryfs/pkg/role/dir"
	"fmt"

	"github.com/google/uuid"
)

/*
	HostManager is responsible for managing all hosts in the cluster
*/
type HostManager struct {
	Hosts   []*Host
	hostMap map[string]*Host
}

func (hostMg *HostManager) New() {
	hostMg.initHostMap()
}

func (hostMg *HostManager) PrintHostMap() {
	fmt.Println(hostMg.hostMap)
}

func (hostMg *HostManager) GetHostByHostId(HostId string) (Host, error) {
	if host, ok := hostMg.hostMap[HostId]; ok {
		return *host, nil
	}
	return Host{}, fmt.Errorf("NotFound")
}

func (hostMg *HostManager) GetHostPointerByHostId(HostId string) (*Host, error) {
	if host, ok := hostMg.hostMap[HostId]; ok {
		return host, nil
	}
	return nil, fmt.Errorf("NotFound")
}

func (hostMg *HostManager) initHostMap() error {
	hostMg.hostMap = make(map[string]*Host)
	for _, host := range hostMg.Hosts {
		hostMg.hostMap[host.HostId] = host
	}

	return nil
}

func (hostMg *HostManager) InitAllHosts(configHosts []ConfigHost, dirManager *dir.DirManager) error {
	hostMg.Hosts = make([]*Host, 0)
	for _, configHost := range configHosts {
		fmt.Printf("cfg host: %s\n", configHost.Address)
		dirIds := make([]string, 0)
		hostId := uuid.New().String()
		for _, d := range configHost.Dirs {
			fmt.Printf("dir: %s\n", d)
			id, _ := dirManager.CreateDir(d, hostId)
			dirIds = append(dirIds, id)
		}

		hostMg.Hosts = append(hostMg.Hosts, &Host{
			HostId:   hostId,
			Hostname: configHost.Hostname,
			Address:  configHost.Address,
			Dirs:     dirIds,
		})
	}
	hostMg.New()
	dirManager.InitDirMap()

	return nil
}

func (hostMg *HostManager) AddHost(hostId, hostAddr string, hostDirs []dir.Dir, dirManager *dir.DirManager) (string, error) {
	host, err := hostMg.InitNewHost(hostId, hostAddr, hostDirs, dirManager)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	hostMg.Hosts = append(hostMg.Hosts, host)
	hostMg.hostMap[host.HostId] = host
	return host.HostId, nil
}

func (hostMg *HostManager) InitNewHost(hostId, hostAddr string, hostDirs []dir.Dir, dirManager *dir.DirManager) (*Host, error) {
	dirs := make([]string, 0)
	for _, hostDir := range hostDirs {
		id, _ := dirManager.CreateDir(hostDir.Path, hostId)
		dirManager.SetTotalSpace(id, hostDir.TotalSpace)
		dirManager.SetUsedSpace(id, hostDir.UsedSpace)
		dirs = append(dirs, id)
	}

	host := Host{
		HostId:    hostId,
		Address:   hostAddr,
		Dirs:      dirs,
		HostState: HEALTHY,
	}

	return &host, nil
}
