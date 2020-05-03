package main

import (
	"crypto/sha256"
	"fmt"
	"encoding/binary"
)

type Test struct {
	Num int
	//Host mgt.Host
}

func main()  {
	//config := mgt.LoadConfig()
	//hosts := mgt.LoadHosts(config.Hosts)
	//subgroups := subgrouper.InitSubgroups(hosts)
	//
	//for _, group := range subgroups {
	//	var s = ""
	//	for _, host := range group.Hosts {
	//		s += host.Hostname + ", "
	//	}
	//	fmt.Printf("%s \n", s)
	//}
	s := []byte("hello")
	h := sha256.New()
	h.Write(s)

	bs := h.Sum(nil)
	fmt.Println(binary.BigEndian.Uint64(bs))

}
