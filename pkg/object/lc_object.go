package object

import (
	"bytes"
	"os"
	"path"
	"fmt"
	"cherryfs/internal/etcd"
)

type LocalObject struct {
	Name string
	Size int64
	Hash string
	Path string
	HostId string
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

func (lcObject *LocalObject) PostStore(client etcd.EtcdClient) error {
	putKey := fmt.Sprintf("%s/%s/%s", ObjectKeyPrefix, lcObject.HostId, lcObject.Name)

	//ctx.EtcdCli.CreateEtcdClient(os.Getenv("ETCDADDR"))
	err := client.Put(putKey, "1")

	fmt.Printf("mark %s as put\n", putKey)
	return err
}