package internal

import (
	"path/filepath"
	"runtime/debug"
	"strings"
)

// module path which will be used as the prefix to be trimmed.
var modulePath string

func init() {
	info, ok := debug.ReadBuildInfo()
	if ok {
		// initialize module path.
		modulePath = info.Main.Path
	}
}

// FormatFile returns last two path elements if file path is absolute path,
// otherwise return the path which relative to the module path.
//
// Reference: https://github.com/qbox/kodo/blob/develop/libs/log.v1/logext.go#L117-L135
func FormatFile(file string, lastN int) string {
	if !filepath.IsAbs(file) {
		return strings.TrimPrefix(file, modulePath+"/")
	}

	var lastNthSlash = -1
	var path = file
	for lastN > 0 {
		lastN--
		if pos := strings.LastIndex(path, "/"); pos >= 0 {
			lastNthSlash = pos
			path = path[:pos]
		} else {
			lastNthSlash = -1
			break
		}
	}
	return file[lastNthSlash+1:]
}
