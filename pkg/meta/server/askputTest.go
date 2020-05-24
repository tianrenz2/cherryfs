package server

//
//func AskPutTest() {
//	var testName string = "obj"
//	var testSize int64 = 12
//	var testHash string = "961f8fe103d6b51f8853deb27e7c26a1b6c5057e4da4db44175ad3f1e6aaa1c7"
//	apr := pb.AskPutRequest.AskPutRequest{Name:&testName, Size:&testSize, ObjectHash:&testHash}
//
//	targets, err := AskPut(&apr)
//
//	if err != nil {
//		log.Fatalf("%v\n", err)
//	}
//
//	for _, target := range targets {
//		fmt.Printf("Assign object %s to host: %s's dir %s, dir size: %v\n",
//			testName, target.Host.Address, target.Dir.Path, target.Dir.TotalSpace)
//	}
//}

//func main()  {
//	//GlobalCtx = meta.Startup()
//	GlobalCtx = initialize.LoadClusterConfig()
//	AskPutTest()
//}