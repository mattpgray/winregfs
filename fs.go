package winregfs

import (
	"io/fs"
	"sort"

	"golang.org/x/sys/windows/registry"
)

func FS(key registry.Key, name string) fs.FS {
	return &regFS{key, name}
}

type regFS struct {
	key  registry.Key
	name string
}

var _ fs.FS = (*regFS)(nil)
var _ fs.ReadDirFS = (*regFS)(nil)

func (rFS *regFS) Open(name string) (fs.File, error) {
	return nil, nil
}

func (rFS *regFS) ReadDir(name string) ([]fs.DirEntry, error) {
	key, err := registry.OpenKey(rFS.key, name, registry.READ)
	if err != nil {
		return nil, err
	} // if
	defer key.Close()

	return readDir(key, name)
}

type regDirEntry struct {
	parentKey registry.Key
	name      string
	keyName   string
	mode      fs.FileMode
}

var _ fs.DirEntry = (*regDirEntry)(nil)

func (de *regDirEntry) Name() string {
	return de.keyName
}

func (de *regDirEntry) IsDir() bool {
	return de.mode.IsDir()
}

func (de *regDirEntry) Type() fs.FileMode {
	return de.mode
}

func (de *regDirEntry) Info() (fs.FileInfo, error) {

}

func readDir(key registry.Key, name string) ([]fs.DirEntry, error) {
	subKeys, err := key.ReadSubKeyNames(-1)
	if err != nil {
		return nil, err
	} // if

	values, err := key.ReadValueNames(-1)
	if err != nil {
		return nil, err
	} // if

	dirEntries := make([]fs.DirEntry, 0, len(subKeys)+len(values))

	for _, subKey := range subKeys {
		dirEntries = append(dirEntries, &regDirEntry{key, name, subKey, fs.ModeDir})
	}

	for _, value := range values {
		dirEntries = append(dirEntries, &regDirEntry{key, name, value, fs.ModeDir})
	}

	sort.Slice(dirEntries, func(i, j int) bool { return dirEntries[i].Name() < dirEntries[j].Name() })
	return dirEntries, nil
}
