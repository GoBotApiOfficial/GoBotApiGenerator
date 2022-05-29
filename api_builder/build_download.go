package api_builder

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/api_grabber/types"
	"BotApiCompiler/consts"
	"fmt"
	"golang.org/x/exp/slices"
	"os"
	"path"
	"strings"
)

func (ctx *Context) BuildDownload(listElements map[string]*types.ApiTypeTL) {
	var filesInput []string
	for _, typeScheme := range listElements {
		for _, field := range typeScheme.GetFields() {
			if field.Name == "file_id" {
				filesInput = append(filesInput, typeScheme.GetName())
			}
		}
	}
	outputFileFolder := path.Join(consts.OutputFolder, "download_media.go")
	builder := component.NewBuilder()
	builder.SetPackage(utils.MainPackage())
	builder.AddImport("", fmt.Sprintf("%s/types", consts.PackageName))
	builder.AddImport("", "errors")
	builder.AddFunc(
		"ctx *Client",
		"DownloadMedia",
		[]string{"message types.Message, filePath string"},
		"error",
	)
	for _, method := range ctx.ApiTL.Types["Message"].GetFields() {
		_, fixGeneric := utils.FixArray(method.Types[0])
		isArray := strings.Contains(method.Types[0], "Array")
		if slices.Contains(filesInput, fixGeneric) {
			if isArray {
				builder.AddIf(fmt.Sprintf("len(message.%s) > 0", utils.PrettifyField(method.Name)))
				builder.DeclareVar("bestQuality", fmt.Sprintf("types.%s", fixGeneric)).AddLine()
				builder.AddFor(fmt.Sprintf("_, file := range message.%s", utils.PrettifyField(method.Name)))
				builder.AddIf("file.Width > bestQuality.Width")
				builder.SetVarValue("bestQuality", "file").AddLine()
				builder.CloseBracket().CloseBracket()
				builder.AddReturn("ctx.DownloadFile(bestQuality.FileID, filePath)").AddLine()
				builder.CloseBracket()
			} else {
				builder.AddIf(fmt.Sprintf("message.%s != nil", utils.PrettifyField(method.Name)))
				builder.AddReturn(fmt.Sprintf("ctx.DownloadFile(message.%s.FileID, filePath)", utils.PrettifyField(method.Name))).AddLine()
				builder.CloseBracket()
			}
		}
	}
	builder.AddReturn("errors.New(\"no files found\")").AddLine()
	builder.CloseBracket()
	_ = os.WriteFile(outputFileFolder, builder.Build(), 0755)
}
