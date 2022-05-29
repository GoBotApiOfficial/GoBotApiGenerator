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
)

func (ctx *Context) BuildDownload(listElements map[string]*types.ApiTypeTL) {
	outputFileFolder := path.Join(consts.OutputFolder, "download_message.go")
	builder := component.NewBuilder()
	builder.SetPackage(utils.MainPackage())
	builder.AddImport("", fmt.Sprintf("%s/types", consts.PackageName))
	builder.AddFunc(
		"ctx *Client",
		"DownloadMessage",
		[]string{"message types.Message", "filePath string"},
		"error",
	)
	var filesInput []string
	for _, typeScheme := range listElements {
		for _, field := range typeScheme.GetFields() {
			if field.Name == "file_id" {
				filesInput = append(filesInput, typeScheme.GetName())
			}
		}
	}
	for _, method := range ctx.ApiTL.Types["Message"].GetFields() {
		if slices.Contains(filesInput, utils.PrettifyField(method.Name)) {
			builder.AddIf(fmt.Sprintf("message.%s != nil", utils.PrettifyField(method.Name)))
			builder.AddReturn(fmt.Sprintf("ctx.DownloadFile(message.%s.FileID, filePath)", utils.PrettifyField(method.Name))).AddLine()
			builder.CloseBracket()
		}
	}
	builder.AddReturn("nil").AddLine()
	builder.CloseBracket()
	_ = os.WriteFile(outputFileFolder, builder.Build(), 0755)
}
