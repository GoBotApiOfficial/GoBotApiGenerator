package utils

import (
	"BotApiCompiler/consts"
	"path"
)

func MainPackage() string {
	return path.Base(consts.PackageName)
}
