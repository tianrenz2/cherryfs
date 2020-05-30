package etcd

import (
	"testing"
	"fmt"
)

func Test_PutGetPrefix(t *testing.T) {

}

func Test_PutGet(t *testing.T)  {
	client := EtcdClient{}
	e := client.CreateEtcdClient("127.0.0.1:2380")

	if e != nil {
		t.Error("Create Client: not passed")
	}else {
		t.Log("Create Client: passed")
	}


	var testKey = "test_key"
	var testVal = "test_val"

	err := client.Put(testKey, testVal)

	if err != nil {
		t.Error("Put: not passed")
	}else {
		t.Log("Put: passed")
	}

	val, err := client.Get(testKey)

	if err != nil {
		t.Error("Get: not passed")
	}

	if val != testVal {
		t.Error(fmt.Sprintf("Get: not passed, %s != %s", testVal, val))
	}
}

