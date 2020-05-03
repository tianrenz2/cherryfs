package dir

type Dir struct {
	Uuid string
	Path string
	TotalSpace int64
	UsedSpace int64
	Manager DirManager
}

func (dir *Dir) GetBaseScore() float64 {
	reliefNum := int64(dir.Manager.ReliefNum)
	numerator := dir.TotalSpace - dir.UsedSpace
	denominator := (dir.UsedSpace + reliefNum) / (dir.TotalSpace + reliefNum)
	return float64(numerator / denominator)
}