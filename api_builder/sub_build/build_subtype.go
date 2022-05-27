package sub_build

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/interfaces"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/api_grabber/types"
	"fmt"
	"sort"
)

func BuildSubtype[Scheme interfaces.SchemeInterface](typeScheme Scheme, builder *component.Context, listElements map[string]*types.ApiTypeTL) {
	isMethod := typeScheme.GetType() == "methods"
	commonTypes := make(map[string]*types.CommonTypeTL)
	for _, field := range typeScheme.GetSubTypes() {
		if fieldScheme, ok := listElements[field]; ok {
			for _, field := range fieldScheme.GetFields() {
				if commonTypes[field.Name] == nil {
					commonTypes[field.Name] = &types.CommonTypeTL{
						Count: 1,
						Field: field,
					}
				} else {
					commonTypes[field.Name].Count++
				}
			}
		}
	}
	commonTypesOrdered := make([]types.CommonTypeTL, 0, len(commonTypes))
	for _, field := range commonTypes {
		commonTypesOrdered = append(commonTypesOrdered, *field)
	}
	sort.Slice(commonTypesOrdered, func(i, j int) bool {
		return commonTypesOrdered[i].Field.Name < commonTypesOrdered[j].Field.Name
	})
	builder.InitStruct(typeScheme.GetName())
	for _, field := range commonTypesOrdered {
		jsonName := field.Field.Name
		builder.AddField(
			field.Field.Name,
			utils.FixGeneric(
				field.Field.Optional,
				field.Field.Name,
				field.Field.Types,
				isMethod,
				true,
			),
			jsonName,
		)
	}
	builder.CloseBracket().AddLine()
	builder.AddFunc(fmt.Sprintf("x %s", typeScheme.GetName()), "Kind", nil, "int")
	builder.InitSwitch(fmt.Sprintf("x.%s", utils.PrettifyField(typeScheme.GetTypeIds().CommonName)))
	for _, field := range typeScheme.GetSubTypes() {
		builder.AddCase(true, listElements[field].TypeIds.TypeIds)
		builder.AddReturn(fmt.Sprintf("Type%s", field)).CloseCase().AddLine()
	}
	builder.AddDefault().AddReturn("-1").CloseCase().AddLine()
	builder.CloseBracket().CloseBracket()
}
