package etcd

import (
	"time"
	"go.etcd.io/etcd/clientv3"
	"fmt"
	"context"
)


type EtcdClient struct {
	addr string
	client *clientv3.Client
	cfg clientv3.Config
	ctx context.Context
	timeout time.Duration
}

func (client *EtcdClient)CreateEtcdClient(addr string) (err error) {

	serviceAddr := fmt.Sprintf("http://%s", addr)
	cfg := clientv3.Config{
		Endpoints: []string{serviceAddr},
		DialTimeout: 5 * time.Second,
	}

	client.addr = addr
	client.cfg = cfg
	client.client, err = clientv3.New(cfg)
	client.timeout = 5 * time.Second
	ctx, _ := context.WithTimeout(context.Background(), client.timeout)
	client.ctx = ctx

	return
}

func (client *EtcdClient) Put(key, val string) (error) {
	_, err := client.client.Put(client.ctx, key, val)
	return err
}

func (client *EtcdClient) GetWithPrefix(keyPrefix string) (map[string]string, error){

	resp, err := client.client.Get(client.ctx, keyPrefix, clientv3.WithPrefix())

	var res = make(map[string]string)
	//fmt.Println(resp)
	if err == nil{
		//val := string(resp.Kvs)
		for _, kv := range resp.Kvs {
			res[string(kv.Key)] = string(kv.Value)
		}
	}
	return res, nil
}

func (client *EtcdClient)Get(key string) (string, error) {
	resp, err := client.client.Get(client.ctx, key)

	if len(resp.Kvs) == 0 {
		return "", fmt.Errorf("value does not exist")
	}

	var val = ""
	if err == nil{
		val = string(resp.Kvs[0].Value)
	}

	return val, err
}

func (client *EtcdClient)AmILeader() (bool) {
	myEndpoint := client.client.Endpoints()[0]
	status, _ := client.client.Status(client.ctx, myEndpoint)
	return (*status).Header.MemberId == (*status).Leader
}

func (client *EtcdClient)WatchKey(isPrefix bool, keyword string) clientv3.WatchChan {
	opts := []clientv3.OpOption{}
	if isPrefix {
		opts = append(opts, clientv3.WithPrefix())
	}
	fmt.Println(opts)
	watchChan := client.client.Watch(context.Background(), keyword, opts...)

	return watchChan
}


func main() {
	//cfg := clientv3.Config{
	//	Endpoints: []string{"http://127.0.0.1:22379"},
	//	DialTimeout: 5 * time.Second,
	//}
	//cli, err := clientv3.New(cfg)
	//
	//if err != nil {
	//	fmt.Println("Failed to create client: " + err.Error())
	//}

}