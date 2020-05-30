package chunkmanage

import (
	"syscall"
)

type DirStatus struct {
	All  int64 `json:"all"`
	Used int64 `json:"used"`
	Free int64 `json:"free"`
}

func (chunkCtx *ChunkContext) CollectDirStats()  {
	for i, dir := range chunkCtx.LcDirs {
		dirStatus := DirUsage(dir.Path)
		chunkCtx.LcDirs[i].TotalSpace = dirStatus.All
		chunkCtx.LcDirs[i].UsedSpace = dirStatus.Used
	}
	return
}

func DirUsage(path string) (disk DirStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = int64(fs.Blocks) * int64(fs.Bsize)
	disk.Free = int64(fs.Bfree) * int64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}

