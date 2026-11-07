package utils

import (
	"BotApiCompiler/consts"
	"os"
	"path"
)

func InitializeFolders() {
	consts.OutputFolder = MainPackage()
	if _, err := os.Stat(consts.OutputFolder); os.IsNotExist(err) {
		_ = os.Mkdir(consts.OutputFolder, consts.FolderPermission)
	} else {
		_ = os.RemoveAll(consts.OutputFolder)
		_ = os.Mkdir(consts.OutputFolder, consts.FolderPermission)
	}
	var folders = []string{
		"types/raw",
		"methods",
		"filters",
		"utils",
	}
	for _, folder := range folders {
		if _, err := os.Stat(path.Join(consts.OutputFolder, folder)); os.IsNotExist(err) {
			_ = os.MkdirAll(path.Join(consts.OutputFolder, folder), consts.FolderPermission)
		}
	}
}
