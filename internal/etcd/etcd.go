package etcd

import (
	"time"
	"go.etcd.io/etcd/clientv3"
	"fmt"
	"context"
	"strings"
	"log"
)


type EtcdClient struct {
	addrs   []string
	Client  *clientv3.Client
	cfg     clientv3.Config
	timeout time.Duration
}

func (client *EtcdClient)CreateEtcdClient(addrs string) (err error) {
	addrsList := strings.Split(addrs, ",")
	for index, addr := range addrsList {
		addrsList[index] = fmt.Sprintf("http://%s", addr)
	}
	cfg := clientv3.Config{
		Endpoints: addrsList,
		DialTimeout: 5 * time.Second,
	}

	client.addrs = addrsList
	client.cfg = cfg
	client.Client, err = clientv3.New(cfg)
	client.timeout = 5 * time.Second
	return
}

func (client *EtcdClient) Put(key, val string) (error) {
	ctx, _ := context.WithTimeout(context.Background(), client.timeout)
	_, err := client.Client.Put(ctx, key, val)
	return err
}

func (client *EtcdClient) GetWithPrefix(keyPrefix string) (map[string]string, error){
	ctx, _ := context.WithTimeout(context.Background(), client.timeout)
	resp, err := client.Client.Get(ctx, keyPrefix, clientv3.WithPrefix())

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

func (client *EtcdClient) GetWithSuffix(keySuffix string) (map[string]string, error){
	ctx, _ := context.WithTimeout(context.Background(), client.timeout)
	resp, err := client.Client.Get(ctx, keySuffix, clientv3.WithRange(keySuffix))
	fmt.Printf("%v\n", resp.Kvs)
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
	ctx, _ := context.WithTimeout(context.Background(), client.timeout)
	resp, err := client.Client.Get(ctx, key)
	if len(resp.Kvs) == 0 {
		return "", fmt.Errorf("NoExist")
	}

	var val = ""
	if err == nil{
		val = string(resp.Kvs[0].Value)
	}

	return val, err
}

func (client *EtcdClient)AmILeader() (bool) {
	myEndpoint := client.Client.Endpoints()[0]
	ctx, _ := context.WithTimeout(context.Background(), client.timeout)
	status, _ := client.Client.Status(ctx, myEndpoint)
	return (*status).Header.MemberId == (*status).Leader
}

func (client *EtcdClient)WatchKey(isPrefix bool, keyword string) clientv3.WatchChan {
	opts := []clientv3.OpOption{}
	if isPrefix {
		opts = append(opts, clientv3.WithPrefix())
	}
	fmt.Println(opts)
	watchChan := client.Client.Watch(context.Background(), keyword, opts...)

	return watchChan
}

func (client *EtcdClient) MaintainKey(key string, ttl int64) (<-chan *clientv3.LeaseKeepAliveResponse, error){
	fmt.Printf("maintain %s\n", key)
	// minimum lease TTL is 5-second

	resp, err := client.Client.Grant(context.TODO(), ttl)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Printf("lease id %v\n", resp.ID)
	_, err = client.Client.Put(context.TODO(), key, string("1"), clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	leaseResp, err := client.Client.KeepAlive(context.TODO(), resp.ID)

	return leaseResp, err
}


func main() {
	//cfg := clientv3.Config{
	//	Endpoints: []string{"http://127.0.0.1:22379"},
	//	DialTimeout: 5 * time.Second,
	//}
	cli := EtcdClient{}

	cli.CreateEtcdClient("127.0.0.1:2379,127.0.0.1:22379,127.0.0.1:32379")

	cli.Put("abc/bbc", "b")

	val, _ := cli.GetWithSuffix("bbc")
	fmt.Println(val)
}