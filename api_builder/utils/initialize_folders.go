package utils

import (
	"BotApiCompiler/consts"
	"os"
	"path"
)

func InitializeFolders() {
	consts.OutputFolder = MainPackage()
	if _, err := os.Stat(consts.OutputFolder); os.IsNotExist(err) {
		_ = os.Mkdir(consts.OutputFolder, 0755)
	} else {
		_ = os.RemoveAll(consts.OutputFolder)
		_ = os.Mkdir(consts.OutputFolder, 0755)
	}
	var folders = []string{
		"types/raw",
		"methods",
	}
	for _, folder := range folders {
		if _, err := os.Stat(path.Join(consts.OutputFolder, folder)); os.IsNotExist(err) {
			_ = os.MkdirAll(path.Join(consts.OutputFolder, folder), 0755)
		}
	}
}
