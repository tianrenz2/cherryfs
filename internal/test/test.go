package main

import (
	"cherryfs/pkg/client"
	"os"
	"fmt"
	"log"
)

func main()  {
	cli := client.Client{}
	cli.Init("127.0.0.1:50051")

	if len(os.Args) < 3 {
		log.Fatalf("Not enough arguments")
		return
	}

	ObjectKey := os.Args[1]
	ObjectPath := os.Args[2]

	err := cli.Put(ObjectKey, ObjectPath)
	//err := cli.Get(ObjectKey, ObjectPath)

	if err != nil {
		fmt.Errorf("failed to put object: %v\n", err)
	}
}

func Get()  {
	
}