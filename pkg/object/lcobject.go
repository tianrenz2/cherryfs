package object

import (
	"bytes"
	"os"
	"path"
	"fmt"
	"cherryfs/pkg/chunk/chunkmanage"
)

type LocalObject struct {
	Name string
	Size int64
	Hash string
	Path string
}

func (lcObject *LocalObject) ObjectStore(data bytes.Buffer) (error) {
	destPath := path.Join(lcObject.Path, lcObject.Name)
	fmt.Println(destPath)
	file, err := os.Create(destPath)
	_, err = data.WriteTo(file)

	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

func (lcObject *LocalObject) PostStore(ctx chunkmanage.ChunkContext) error {

	putKey := ObjectKeyPrefix + lcObject.Name + "/" + ctx.HostId

	ctx.EtcdCli.CreateEtcdClient(os.Getenv("ETCD_ADDR"))
	err := ctx.EtcdCli.Put(putKey, "1")

	fmt.Printf("put %s\n", putKey)

	fmt.Println(putKey)
	return err
}