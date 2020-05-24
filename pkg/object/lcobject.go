package object

import (
	"bytes"
	"os"
	"path"
	"fmt"
	"cherryfs/pkg/context"
)

type LocalObject struct {
	Name string
	Size int64
	Hash string
	Path string
}

func (lcObject *LocalObject) ObjectStore(data bytes.Buffer) (error) {
	file, err := os.Create(path.Join(lcObject.Path, lcObject.Name))
	_, err = data.WriteTo(file)

	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

func (lcObject *LocalObject) PostStore(ctx context.Context) error {

	putKey := ObjectKeyPrefix + lcObject.Name + "/" + ctx.HostId

	ctx.EtcdCli.CreateEtcdClient(os.Getenv("ETCD_ADDR"))
	err := ctx.EtcdCli.Put(putKey, "1")

	fmt.Println(putKey)
	return err
}