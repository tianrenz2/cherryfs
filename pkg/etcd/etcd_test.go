package etcd

import (
	"testing"
	"fmt"
)


func Test_PutGetPrefix(t *testing.T) {
	client := EtcdClient{}
	e := client.CreateEtcdClient("127.0.0.1:2380")

	var prefix = "test_key"

	if e != nil {
		t.Error("Create Client: not passed")
	}else {
		t.Log("Create Client: passed")
	}

	var testKey1 = prefix + "1"
	var testVal1 = "test_val1"

	err := client.Put(testKey1, testVal1)

	if err != nil {
		t.Error("Put: not passed")
	}else {
		t.Log("Put: passed")
	}

	var testKey2 = prefix + "2"
	var testVal2 = "test_val2"

	err = client.Put(testKey2, testVal2)
	if err != nil {
		t.Error("Put: not passed")
	}else {
		t.Log("Put: passed")
	}

	res, _ := client.GetWithPrefix(prefix)

	fmt.Println(res)

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

