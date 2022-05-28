package api_builder

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/interfaces"
	"BotApiCompiler/api_builder/sub_build"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/api_grabber/types"
	"BotApiCompiler/consts"
	"fmt"
	"os"
	"path"
	"strings"
)

func BuildComponent[Scheme interfaces.SchemeInterface](typeScheme Scheme, listElements map[string]*types.ApiTypeTL, listConsts []string) {
	typeName := typeScheme.GetType()
	isMethod := strings.Contains(typeName, "methods")
	var outputFileFolder string
	if isMethod {
		outputFileFolder = path.Join(consts.OutputFolder, typeName, fmt.Sprintf("%s.go", utils.FixName(typeScheme.GetName())))
	} else {
		outputFileFolder = path.Join(consts.OutputFolder, typeName, fmt.Sprintf("%s.go", utils.FixName(typeScheme.GetName())))
	}
	builder := component.NewBuilder()
	builder.SetPackage(typeName)
	builder.AddDocumentation(utils.FixStructName(typeScheme.GetName()), typeScheme.GetDescription())
	if len(typeScheme.GetSubTypes()) > 0 && !typeScheme.IsSendMethod() {
		sub_build.BuildSubtype(typeScheme, &builder, listElements)
	} else {
		for _, field := range listElements {
			for _, fieldType := range field.GetSubTypes() {
				if typeScheme.GetName() == fieldType && len(field.GetSubTypes()) > 0 && !field.IsSendMethod() {
					return
				}
			}
		}
		inputFiles := sub_build.BuildType(typeScheme, &builder, listElements)
		sub_build.BuildFiles(typeScheme, &builder, inputFiles)
		sub_build.BuildMarshaller(typeScheme, &builder, listElements)
		sub_build.BuildMethodName(typeScheme, &builder)
		sub_build.BuildParser(typeScheme, &builder, listConsts)
	}
	if builder.HaveBody() {
		_ = os.WriteFile(outputFileFolder, builder.Build(), 0755)
	}
}
