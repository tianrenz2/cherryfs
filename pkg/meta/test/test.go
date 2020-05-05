package main

import (
	"fmt"
	"encoding/json"
	"encoding/base64"
)

type Test struct {
	Num int

	//Host mgt.Host
}

func main()  {
	//config := mgt.LoadConfig()
	//hosts := mgt.LoadHosts(config.Hosts)
	//subgroups := subgroup.InitSubgroups(hosts)
	//
	//for _, group := range subgroups {
	//	var s = ""
	//	for _, host := range group.Hosts {
	//		s += host.Hostname + ", "
	//	}
	//	fmt.Printf("%s \n", s)
	//}
	test := Test{Num:4}
	fmt.Println(test)
	m, _ := json.Marshal(test)

	fmt.Println(m)

	g := base64.StdEncoding.EncodeToString(m)

	var b []byte

	b, _ = base64.StdEncoding.DecodeString(g)

	fmt.Println(b)

	var t Test

	json.Unmarshal(b, &t)

	fmt.Println(t)

	//s := []byte("hello")
	//h := sha256.New()
	//h.Write(s)
	//
	//bs := h.Sum(nil)
	//fmt.Println(binary.BigEndian.Uint64(bs))

}
