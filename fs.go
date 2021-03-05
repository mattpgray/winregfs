package winregfs

import (
	"golang.org/x/sys/windows/registry"
)

type FS struct {
	key registry.Key
}
