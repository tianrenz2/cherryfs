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
	HostpKeyPrefix = "role/host"
	DirKeyPrefix = "role/dir"
)

type Context struct {
	SGManager subgroup.SubGroupManager
	HManager host.HostManager
	DManager dir.DirManager
	EtcdCli etcd.EtcdClient
}


func (ctx *Context) PersistClusterConfig() (error) {
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

func (ctx *Context) DecodeSubgroups() (error) {
	sgmap, err := ctx.EtcdCli.GetWithPrefix(SubGroupKeyPrefix)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

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