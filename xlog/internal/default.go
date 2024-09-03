package internal

import (
	"io"
	"os"
	"sync"
)

var (
	DefaultW      io.Writer  = os.Stdout
	DefaultWMutex sync.Mutex = sync.Mutex{}

	DefaultLevel     Level  = LevelInfo
	DefaultFormat    Format = FormatClassical
	DefaultAddSource bool   = true
	DefaultAutoReqID bool   = true
)
