package winregfs

import (
	"errors"
	"io/fs"

	"golang.org/x/sys/windows/registry"
)

type regFile struct {
	name string
	key  registry.Key
}

var _ fs.File = (*regFile)(nil)

func (f *regFile) Stat() (fs.FileInfo, error) {
	return FileInfo(f.name, f.key)
}

var ErrUnsupported = errors.New("Read method unsupported for Window registry")

func (f *regFile) Read([]byte) (int, error) {
	return 0, ErrUnsupported
}

func (f *regFile) Close() error {
	return f.key.Close()
}
