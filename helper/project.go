package helper

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	rootpath   = filepath.Dir(filepath.Dir(b))
)

func ProjectRootPath() string {
	return rootpath
}
