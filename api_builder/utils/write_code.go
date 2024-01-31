package utils

import (
	"BotApiCompiler/consts"
	"go/format"
	"os"
)

func WriteCode(outputFileFolder string, output []byte) {
	source, err := format.Source(output)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(outputFileFolder, source, consts.FolderPermission)
	if err != nil {
		panic(err)
	}
}
