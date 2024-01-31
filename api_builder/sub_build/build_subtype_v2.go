package sub_build

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/interfaces"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/api_grabber/types"
	"fmt"
	"golang.org/x/exp/slices"
)

func BuildSubtypeV2[Scheme interfaces.SchemeInterface](typeScheme Scheme, builder *component.Context, listElements map[string]*types.ApiTypeTL) {
	isMethod := typeScheme.GetType() == "methods"
	if !isMethod {
		var typeV2 *types.ApiTypeTL
		for _, field := range listElements {
			for _, fieldType := range field.GetSubTypes() {
				if typeScheme.GetName() == fieldType && len(field.GetSubTypes()) > 0 &&
					!field.IsSendMethod() && typeScheme.GetTypeIds() == nil {
					typeV2 = field
				}
			}
		}
		if typeV2 != nil {
			builder.AddLine()
			builder.AddFunc(fmt.Sprintf("x %s", typeScheme.GetName()), "Kind", nil, "int")
			builder.InitSwitch(fmt.Sprintf("x.%s", utils.PrettifyField(typeV2.GetTypeIds().CommonName)))
			for _, field := range typeV2.GetSubTypes() {
				if listElements[field].TypeIds != nil {
					isString := false
					for _, el := range listElements[field].GetFields() {
						if el.Name == typeV2.GetTypeIds().CommonName {
							isString = slices.Contains(el.Types, "String")
						}
					}
					builder.AddCase(isString, listElements[field].TypeIds.TypeIds)
					builder.AddReturn(fmt.Sprintf("Type%s", field)).CloseCase().AddLine()
				}
			}
			for _, field := range typeV2.GetSubTypes() {
				if listElements[field].TypeIds == nil {
					builder.AddDefault().AddReturn(fmt.Sprintf("Type%s", field)).CloseCase().AddLine()
					break
				}
			}
			builder.CloseBracket().CloseBracket()
		}
	}
}
