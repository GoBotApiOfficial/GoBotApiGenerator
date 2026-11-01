package utils

import (
	"BotApiCompiler/consts"
	"fmt"
	"os"
	"path"
)

func InitializeMod() {
	goModFile := path.Join(consts.OutputFolder, "go.mod")
	body := fmt.Sprintf("module %s\n\ngo 1.23.0\n\nrequire github.com/Laky-64/gologging v1.1.0\n", consts.PackageName)
	_ = os.WriteFile(goModFile, []byte(body), consts.FolderPermission)
}
