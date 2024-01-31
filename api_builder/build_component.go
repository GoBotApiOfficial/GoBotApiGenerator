package api_builder

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/interfaces"
	"BotApiCompiler/api_builder/sub_build"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/api_grabber/types"
	"BotApiCompiler/consts"
	"fmt"
	"path"
)

func BuildComponent[Scheme interfaces.SchemeInterface](typeScheme Scheme, listElements map[string]*types.ApiTypeTL, listConsts []string) {
	typeName := typeScheme.GetType()
	var outputFileFolder string
	if typeScheme.GetName() == "ResponseParameters" {
		outputFileFolder = path.Join(consts.OutputFolder, typeName, "raw", fmt.Sprintf("%s.go", utils.FixName(typeScheme.GetName())))
		typeName = "raw"
	} else {
		outputFileFolder = path.Join(consts.OutputFolder, typeName, fmt.Sprintf("%s.go", utils.FixName(typeScheme.GetName())))
	}
	builder := component.NewBuilder()
	builder.SetPackage(typeName)
	builder.AddDocumentation(utils.FixStructName(typeScheme.GetName()), typeScheme.GetDescription())
	if len(typeScheme.GetSubTypes()) > 0 && !typeScheme.IsSendMethod() && typeScheme.GetTypeIds() != nil {
		for _, field := range typeScheme.GetSubTypes() {
			if listElements[field].TypeIds == nil {
				builder.AddType(utils.FixStructName(typeScheme.GetName()), listElements[field].GetName()).AddLine()
				utils.WriteCode(outputFileFolder, builder.Build())
				return
			}
		}
		sub_build.BuildSubtype(typeScheme, builder, listElements)
	} else {
		for _, field := range listElements {
			for _, fieldType := range field.GetSubTypes() {
				if typeScheme.GetName() == fieldType && len(field.GetSubTypes()) > 0 &&
					!field.IsSendMethod() && typeScheme.GetTypeIds() != nil {
					return
				}
			}
		}
		inputFiles := sub_build.BuildType(typeScheme, builder, listElements)
		sub_build.BuildFiles(typeScheme, builder, inputFiles)
		sub_build.BuildMarshaller(typeScheme, builder, listElements)
		sub_build.BuildMethodName(typeScheme, builder)
		sub_build.BuildParser(typeScheme, builder, listConsts)
		sub_build.BuildSubtypeV2(typeScheme, builder, listElements)
	}
	if builder.HaveBody() {
		utils.WriteCode(outputFileFolder, builder.Build())
	}
}
