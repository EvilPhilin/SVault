package cfg

import (
	"path/filepath"
	"runtime"
)

// Change it if u move cfg directory!
var (
	_, b, _, _ = runtime.Caller(0)
	
    Root = filepath.Join(filepath.Dir(b), "..")
	DataPath = filepath.Join(Root, "/data")
)