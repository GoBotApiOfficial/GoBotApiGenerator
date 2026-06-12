package sub_build

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/interfaces"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/api_grabber/types"
	"fmt"
	"sort"
	"strings"
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
	isPolymorphicValue := utils.IsPolymorphicValue(typeScheme.GetName(), typeScheme.GetDescription())
	builder.InitStruct(typeScheme.GetName())
	var aliasFields []AliasFieldTL
	hasPolyFields := false
	for _, field := range commonTypesOrdered {
		jsonName := field.Field.Name
		fieldTypes := field.Field.Types
		parseFunc := ""
		baseTypeName := strings.TrimPrefix(fieldTypes[0], "Array of ")
		if baseScheme := listElements[baseTypeName]; baseScheme != nil &&
			utils.IsPolymorphicValue(baseScheme.GetName(), baseScheme.GetDescription()) {
			fieldTypes = []string{strings.ReplaceAll(fieldTypes[0], baseTypeName, baseTypeName+"Value")}
			parseFunc = fmt.Sprintf("Parse%sValue", baseTypeName)
			hasPolyFields = true
		}
		genericName := utils.FixGeneric(
			field.Field.Optional,
			field.Field.Name,
			fieldTypes,
			isMethod,
			true,
		)
		if len(parseFunc) > 0 {
			genericName = strings.TrimPrefix(genericName, "*")
		}
		if genericName == typeScheme.GetName() {
			genericName = "*" + genericName
		}
		aliasFields = append(aliasFields, AliasFieldTL{
			Name:      field.Field.Name,
			Generic:   genericName,
			JsonName:  jsonName,
			ParseFunc: parseFunc,
		})
		builder.AddField(
			field.Field.Name,
			genericName,
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
	if isPolymorphicValue {
		BuildPolyValue(builder, typeScheme.GetName())
	}
	if hasPolyFields && !isMethod {
		BuildPolyUnmarshal(builder, typeScheme.GetName(), aliasFields)
	}
}
