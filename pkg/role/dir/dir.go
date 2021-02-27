package dir

type Dir struct {
	DirId      string
	Path       string
	HostId     string
	TotalSpace int64
	UsedSpace  int64
	Manager    DirSubGroupSetting

}

func (dir *Dir) GetBaseScore() float64 {
	reliefNum := int64(dir.Manager.ReliefNum)
	numerator := float64(dir.TotalSpace - dir.UsedSpace)
	denominator := float64(dir.UsedSpace + reliefNum) / float64(dir.TotalSpace + reliefNum)
	return float64(numerator / denominator)
}

func (dir *Dir) UpdateTotalSpace(space int64)  {
	dir.TotalSpace = space
}

//func (dir *Dir) UpdateField(newFields interface{})  {
//	entityType := reflect.TypeOf(newFields).Elem()
//	for i := 0; i < entityType.NumField(); i++ {
//
//	}
//}