package utils

import (
	"BotApiCompiler/consts"
	"fmt"
	"os"
	"path"
)

func InitializeMod() {
	goModFile := path.Join(consts.OutputFolder, "go.mod")
	body := fmt.Sprintf("module %s\n\ngo 1.18", consts.PackageName)
	_ = os.WriteFile(goModFile, []byte(body), consts.FolderPermission)
}
