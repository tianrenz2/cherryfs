package subgrouper

func SubgroupHash(key, subgroupNum, stripe int) int {
	//bitmask := key &
	return key % subgroupNum
}
