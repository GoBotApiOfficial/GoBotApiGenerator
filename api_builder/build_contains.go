package api_builder

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/api_grabber/types"
	"BotApiCompiler/consts"
	"fmt"
	"golang.org/x/exp/slices"
	"path"
	"strings"
)

func (ctx *Context) BuildContains(listElements map[string]*types.ApiTypeTL) {
	var filesInput []string
	for _, typeScheme := range listElements {
		for _, field := range typeScheme.GetFields() {
			if field.Name == "file_id" {
				filesInput = append(filesInput, typeScheme.GetName())
			}
		}
	}
	outputFileFolder := path.Join(consts.OutputFolder, "utils", "contains_files.go")
	builder := component.NewBuilder()
	builder.SetPackage("utils")
	builder.AddImport("", fmt.Sprintf("%s/types", consts.PackageName))
	builder.AddFunc(
		"",
		"ContainsFiles",
		[]string{"message types.Message"},
		"bool",
	)
	var filesCheck []string
	for _, method := range ctx.ApiTL.Types["Message"].GetFields() {
		arr, fixGeneric := utils.FixArray(method.Types[0])
		if slices.Contains(filesInput, fixGeneric) {
			if len(arr) > 0 {
				filesCheck = append(filesCheck, fmt.Sprintf("len(message.%s) > 0", utils.PrettifyField(method.Name)))
			} else {
				filesCheck = append(filesCheck, fmt.Sprintf("message.%s != nil", utils.PrettifyField(method.Name)))
			}
		}
	}
	builder.AddReturn(strings.Join(filesCheck, " || \n")).AddLine()
	builder.CloseBracket()
	utils.WriteCode(outputFileFolder, builder.Build())
}
