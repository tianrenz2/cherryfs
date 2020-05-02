package main

type Dir struct {
	uuid string
	name string
	totalSpace int64
	usedSpace int64
}

type Host struct {
	address string
	dirs []Dir
}

type SubGroup struct {
	hosts []Host
}

type GlobalConfig struct {
	subGroupNum int
	subGroups []SubGroup
}
