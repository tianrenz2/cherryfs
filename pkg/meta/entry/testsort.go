package main

import (
	"sort"
	"strconv"
	"math/rand"
	"fmt"
)

func main()  {
	type TargetSummary struct {
		HostId string
		ObjNum int
	}

	var targetItems []TargetSummary

	for i:=0; i<10; i++ {
		targetItems = append(targetItems, TargetSummary{
			HostId: strconv.Itoa(i),
			ObjNum: rand.Int(),
		})
	}

	fmt.Println(targetItems)

	sort.Slice(targetItems[:], func(i, j int) bool {
		return targetItems[i].ObjNum > targetItems[j].ObjNum
	})

	fmt.Println(targetItems)

}
