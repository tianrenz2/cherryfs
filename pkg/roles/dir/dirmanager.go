package dir

import (
	"fmt"
	"encoding/base64"
)

/*
	Dir Manager is responsible for managing all dirs in the cluster
*/
type DirManager struct {
	Dirs []*Dir
	dirMap map[string]*Dir
}

func (dirMg *DirManager) GetDirByDirId(dirId string) (Dir, error) {
	if dir, ok := dirMg.dirMap[dirId]; ok {
		return *dir, nil
	}

	return Dir{}, fmt.Errorf("could not find dir for %s", dirId)
}

func (dirMg *DirManager) UpdateDirSpaceByDirId(dirId string, space int64) {
	dirMg.dirMap[dirId].TotalSpace = space
}

func (dirMg *DirManager) InitDirMap() (error) {
	dirMg.dirMap = make(map[string]*Dir)
	for _, dir := range dirMg.Dirs {
		dirMg.dirMap[dir.DirId] = dir
	}

	return nil
}

func (dirMg *DirManager) CreateDir(path, hostId string) (dirId string, err error) {
	dirId = dirMg.GenDirId(path, hostId)
	newDir := Dir{Path:path, HostId:hostId, DirId:dirId}
	dirMg.HireDir(&newDir)
	return dirId, nil
}

func (dirMg *DirManager) HireDir(dir *Dir)  {
	dirMg.Dirs = append(dirMg.Dirs, dir)
	dirMg.dirMap[dir.DirId] = dir
}

func (dirMg *DirManager) GenDirId(path, hostId string) string {
	return fmt.Sprintf("%s-%s", hostId, base64.StdEncoding.EncodeToString([]byte(path)))
}

func (dirMg *DirManager) SetUsedSpace(dirId string, space int64) error {
	dirMg.dirMap[dirId].UsedSpace = space
	return nil
}

func (dirMg *DirManager) SetTotalSpace(dirId string, space int64) error {
	dirMg.dirMap[dirId].TotalSpace = space
	return nil
}