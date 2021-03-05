package winregfs

import (
	"io/fs"
	"time"

	"golang.org/x/sys/windows/registry"
)

func FileInfo(name string, key registry.Key) (fs.FileInfo, error) {
	info, err := key.Stat()
	if err != nil {
		return nil, err
	}

	return &regFileInfo{name, key, info}, nil
}

type regFileInfo struct {
	name string
	key  registry.Key
	info *registry.KeyInfo
}

var _ fs.FileInfo = (*regFileInfo)(nil)

func (i *regFileInfo) Name() string {
	return i.name
}

func (i *regFileInfo) Size() int64 {
	return 0 // invalid
}

func (i *regFileInfo) Mode() fs.FileMode {
	return 0
}

func (i *regFileInfo) ModTime() time.Time {
	return i.info.ModTime()
}

func (i *regFileInfo) IsDir() bool {
	return true // keys are always dirs.
}

func (i *regFileInfo) Sys() interface{} {
	return i.key
}
