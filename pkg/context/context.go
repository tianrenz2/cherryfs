package context

import (
	"cherryfs/pkg/roles/host"
	"cherryfs/pkg/roles/dir"
	"cherryfs/pkg/meta/subgroup"
	"cherryfs/pkg/meta/etcd"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"strconv"
)

const (
	SubGroupKeyPrefix = "role/sg"
	HostKeyPrefix     = "role/host"
	DirKeyPrefix      = "role/dir"
)

/*
	Context contains all the manager-level units in the cluster
	It also provides the entry to persist and recover the cluster state
*/
type Context struct {
	SGManager *subgroup.SubGroupManager
	HManager *host.HostManager
	DManager *dir.DirManager
	EtcdCli etcd.EtcdClient
}

func (ctx *Context) InitManagers() {
	ctx.SGManager = &subgroup.SubGroupManager{}
	ctx.HManager = &host.HostManager{}
	ctx.DManager = &dir.DirManager{}
}


func (ctx *Context) PersistCluster() (error) {
	ctx.PersistSubGroups()
	ctx.PersistHosts()
	ctx.PersistDirs()
	return nil
}

func (ctx *Context) PersistSubGroups() (error) {
	for _, sg := range ctx.SGManager.SubGroups {
		encodedStr, err := ctx.EncodeStruct(sg)
		if err != nil {
			return fmt.Errorf("failed to encode subgroup: %d, %v", sg.SubGroupId, err)
		}

		b, _ := ctx.DecodeStruct(encodedStr)
		var s subgroup.SubGroup
		json.Unmarshal(b, &s)

		sgKey := fmt.Sprintf("%s/%s", SubGroupKeyPrefix, strconv.Itoa(sg.SubGroupId))
		err = ctx.EtcdCli.Put(sgKey, encodedStr)
		if err != nil {
			return fmt.Errorf("failed persist subrgoup: %d, %v", sg.SubGroupId, err)
		}
	}
	return nil
}

func (ctx *Context) PersistHosts() (error) {
	for _, h := range ctx.HManager.Hosts {
		encodedStr, err := ctx.EncodeStruct(h)
		if err != nil {
			return fmt.Errorf("failed to encode subgroup: %s, %v", h.HostId, err)
		}

		b, _ := ctx.DecodeStruct(encodedStr)
		var s subgroup.SubGroup
		json.Unmarshal(b, &s)

		sgKey := fmt.Sprintf("%s/%s", HostKeyPrefix, h.HostId)
		err = ctx.EtcdCli.Put(sgKey, encodedStr)
		if err != nil {
			return fmt.Errorf("failed persist host: %d, %v", h.HostId, err)
		}
	}
	return nil
}

func (ctx *Context) PersistDirs() (error) {
	for _, d := range ctx.DManager.Dirs {
		encodedStr, err := ctx.EncodeStruct(d)
		if err != nil {
			return fmt.Errorf("failed to encode dir: %s, %v", d.DirId, err)
		}

		b, _ := ctx.DecodeStruct(encodedStr)
		var s subgroup.SubGroup
		json.Unmarshal(b, &s)

		sgKey := fmt.Sprintf("%s/%s", DirKeyPrefix, d.DirId)
		err = ctx.EtcdCli.Put(sgKey, encodedStr)
		if err != nil {
			return fmt.Errorf("failed persist subrgoup: %s, %v", d.DirId, err)
		}
	}
	return nil
}

func (ctx *Context) RecoverCluster() (error) {
	ctx.DecodeSubgroups()
	ctx.DecodeHosts()
	ctx.DecodeDirs()

	ctx.HManager.InitHostMap()
	ctx.DManager.InitDirMap()

	return nil
}

func (ctx *Context) DecodeSubgroups() (error) {
	sgmap, err := ctx.EtcdCli.GetWithPrefix(SubGroupKeyPrefix)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	ctx.SGManager.SubGroups = make([]subgroup.SubGroup, 0)

	for k, v := range sgmap {
		sgnameList := strings.Split(k, "/")
		sgname := sgnameList[len(sgnameList) - 1]
		var sg subgroup.SubGroup

		sgBytes, err := ctx.DecodeStruct(v)

		err = json.Unmarshal(sgBytes, &sg)
		if err != nil {
			fmt.Errorf("failed to decode subgroup %s: %v", sgname, err)
		} else {
			ctx.SGManager.SubGroups = append(ctx.SGManager.SubGroups, sg)
		}
	}

	return nil
}

func (ctx *Context) DecodeHosts() (error) {
	hostmap, err := ctx.EtcdCli.GetWithPrefix(HostKeyPrefix)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	ctx.HManager.Hosts = make([]*host.Host, 0)

	for k, v := range hostmap {
		hostIdList := strings.Split(k, "/")
		hostId := hostIdList[len(hostIdList) - 1]
		var h host.Host

		dirBytes, err := ctx.DecodeStruct(v)

		err = json.Unmarshal(dirBytes, &h)
		if err != nil {
			fmt.Errorf("failed to decode host %s: %v", hostId, err)
		} else {
			ctx.HManager.Hosts = append(ctx.HManager.Hosts, &h)
		}
	}

	return nil
}

func (ctx *Context) DecodeDirs() (error) {
	dirmap, err := ctx.EtcdCli.GetWithPrefix(DirKeyPrefix)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	ctx.DManager.Dirs = make([]*dir.Dir, 0)

	for k, v := range dirmap {
		dirIdList := strings.Split(k, "/")
		dirId := dirIdList[len(dirIdList) - 1]
		var d dir.Dir

		dirBytes, err := ctx.DecodeStruct(v)

		err = json.Unmarshal(dirBytes, &d)
		if err != nil {
			fmt.Errorf("failed to decode dir %s: %v", dirId, err)
		} else {
			ctx.DManager.Dirs = append(ctx.DManager.Dirs, &d)
		}
	}

	return nil
}

func (ctx *Context) EncodeStruct(structure interface{}) (string, error) {
	bytes, e := json.Marshal(structure)
	if e != nil {
		return string(""), fmt.Errorf("%v", e)
	}
	res := base64.StdEncoding.EncodeToString(bytes)
	return res, nil
}

func (ctx *Context) DecodeStruct(encodedHash string) ([]byte, error) {
	bytes, err := base64.StdEncoding.DecodeString(encodedHash)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return bytes, nil
}