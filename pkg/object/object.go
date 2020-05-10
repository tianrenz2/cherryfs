package object

import (
	"bytes"
	"os"
	"path"
	"fmt"
)

type LocalObject struct {
	Name string
	Size int64
	Hash string
	Path string
}

func (lcObject *LocalObject) ObjectStore(data bytes.Buffer) (error) {
	file, err := os.Create(path.Join(lcObject.Path, lcObject.Name))

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	_, err = data.WriteTo(file)

	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}
