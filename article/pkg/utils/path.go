package utils

import (
	"path"
	"runtime"
	"strings"
)

// ProjectPath 获取项目根目录	[level: 层级, 当前调用的绝对路径向上的目录层级]
func ProjectPath(level int) string {
	_, p, _, ok := runtime.Caller(1)
	if ok {
		p = path.Dir(p)
		var sep = "/"
		var pathAnalysis = strings.Split(p, sep)
		var index = len(pathAnalysis) - level
		p = strings.Join(pathAnalysis[:index], sep)
	}

	return p
}
