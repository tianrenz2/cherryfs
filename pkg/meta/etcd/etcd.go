package etcd

import (
	"time"
	"github.com/coreos/etcd/clientv3"
	"fmt"
	"context"
	"strconv"
)

type EtcdClient struct {
	addr string
	port int
	client *clientv3.Client
	cfg clientv3.Config
	ctx context.Context
	timeout time.Duration
}

func (client *EtcdClient)CreateEtcdClient(addr string, port int) (err error) {

	serviceAddr := fmt.Sprintf("http://%s:%s", addr, strconv.Itoa(port))
	fmt.Println("using " + serviceAddr)
	cfg := clientv3.Config{
		Endpoints: []string{serviceAddr},
		DialTimeout: 5 * time.Second,
	}

	client.addr = addr
	client.port = port
	client.cfg = cfg
	client.client, err = clientv3.New(cfg)
	client.timeout = time.Second
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
	var val = ""
	if err == nil{
		val = string(resp.Kvs[0].Value)
	}

	fmt.Println(val)

	return val, err
}

func main()  {
	cfg := clientv3.Config{
		Endpoints: []string{"http://127.0.0.1:2380"},
		DialTimeout: 5 * time.Second,
	}
	cli, err := clientv3.New(cfg)

	if err != nil {
		fmt.Println("Failed to create client: " + err.Error())
	}

	timeout := time.Second
	ctx, _ := context.WithTimeout(context.Background(), timeout)

	_, err = cli.Put(ctx, "vv", "no")


	if err != nil {
		fmt.Println("Failed to put the key: " + err.Error())
	}

	resp, err := cli.Get(ctx, "vv")

	if err != nil {
		fmt.Println("Fialed to get key: " + err.Error())
	}

	fmt.Println(string(resp.Kvs[0].Value))
}