package main

import (
	"BotApiCompiler/api_builder"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/api_grabber"
	"BotApiCompiler/consts"
	"bytes"
	"embed"
	"os"
	"path"
	"strings"
)

//go:embed templates
var templatesFolder embed.FS

//go:embed VERSION
var version []byte

var botApiVersion string

func main() {
	client := api_builder.Client(api_grabber.Client().DownloadApiTL()).Build()
	botApiVersion = client.ApiTL.Version
	CopyRecursivePath("templates", consts.OutputFolder)
	_ = os.WriteFile(path.Join(consts.OutputFolder, "VERSION"), version, 0755)
}

func CopyRecursivePath(src, dst string) {
	files, _ := templatesFolder.ReadDir(src)
	for _, file := range files {
		if file.IsDir() {
			CopyRecursivePath(path.Join(src, file.Name()), path.Join(dst, file.Name()))
		} else {
			CopyFile(path.Join(src, file.Name()), path.Join(dst, strings.ReplaceAll(file.Name(), path.Ext(file.Name()), ".go")))
		}
	}
}

func CopyFile(src, dst string) {
	if _, err := os.Stat(path.Dir(dst)); os.IsNotExist(err) {
		_ = os.MkdirAll(path.Dir(dst), 0755)
	}
	content, _ := templatesFolder.ReadFile(src)
	content = bytes.ReplaceAll(content, []byte("%PACKAGE%"), []byte(consts.PackageName))
	content = bytes.ReplaceAll(content, []byte("%pkg%"), []byte(utils.MainPackage()))
	content = bytes.ReplaceAll(content, []byte("%BOT_API_VERSION%"), []byte(botApiVersion))
	content = bytes.ReplaceAll(content, []byte("%VERSION%"), version)
	_ = os.WriteFile(dst, content, 0755)
}
