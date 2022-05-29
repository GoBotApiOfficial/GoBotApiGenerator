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

func (ctx *Context) BuildContains(listElements map[string]*types.ApiTypeTL) {
	var filesInput []string
	for _, typeScheme := range listElements {
		for _, field := range typeScheme.GetFields() {
			if field.Name == "file_id" {
				filesInput = append(filesInput, typeScheme.GetName())
			}
		}
	}
	outputFileFolder := path.Join(consts.OutputFolder, "contains_files.go")
	builder := component.NewBuilder()
	builder.SetPackage(utils.MainPackage())
	builder.AddImport("", fmt.Sprintf("%s/types", consts.PackageName))
	builder.AddFunc(
		"",
		"ContainsFiles",
		[]string{"message types.Message"},
		"bool",
	)
	var filesCheck []string
	for _, method := range ctx.ApiTL.Types["Message"].GetFields() {
		if slices.Contains(filesInput, utils.PrettifyField(method.Name)) {
			filesCheck = append(filesCheck, fmt.Sprintf("message.%s != nil", utils.PrettifyField(method.Name)))
		}
	}
	fmt.Println(filesCheck)
	builder.AddReturn(strings.Join(filesCheck, fmt.Sprintf(" || \n%s\t", builder.GetTab()))).AddLine()
	builder.CloseBracket()
	_ = os.WriteFile(outputFileFolder, builder.Build(), 0755)
}
