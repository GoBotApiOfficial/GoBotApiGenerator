package sub_build

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/interfaces"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/api_grabber/types"
	"fmt"
	"golang.org/x/exp/slices"
	"sort"
	"strings"
)

func BuildSubtype[Scheme interfaces.SchemeInterface](typeScheme Scheme, builder *component.Context, listElements map[string]*types.ApiTypeTL) {
	isMethod := typeScheme.GetType() == "methods"
	commonTypes := make(map[string]*types.CommonTypeTL)
	fieldVariants := make(map[string][]string)
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
				if len(field.Types) > 0 && !slices.Contains(fieldVariants[field.Name], field.Types[0]) {
					fieldVariants[field.Name] = append(fieldVariants[field.Name], field.Types[0])
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
	var mixedParsers []objectOrPolyTL
	hasPolyFields := false
	for _, field := range commonTypesOrdered {
		jsonName := field.Field.Name
		fieldTypes := field.Field.Types
		parseFunc := ""
		keepPointer := false
		baseTypeName := strings.TrimPrefix(fieldTypes[0], "Array of ")
		if mixed := findObjectOrPoly(fieldVariants[field.Field.Name], listElements); mixed != nil {
			mixed.FuncName = fmt.Sprintf("parse%s%s", typeScheme.GetName(), utils.PrettifyField(field.Field.Name))
			mixedParsers = append(mixedParsers, *mixed)
			fieldTypes = []string{mixed.ObjectName}
			parseFunc = mixed.FuncName
			keepPointer = true
			hasPolyFields = true
		} else if baseScheme := listElements[baseTypeName]; baseScheme != nil &&
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
		if len(parseFunc) > 0 && !keepPointer {
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
		for _, mixed := range mixedParsers {
			BuildObjectOrPolyParser(builder, mixed.FuncName, mixed.ObjectName, mixed.PolyName, mixed.WrapField)
		}
	}
}

type objectOrPolyTL struct {
	FuncName   string
	ObjectName string
	PolyName   string
	WrapField  string
}

func findObjectOrPoly(variants []string, listElements map[string]*types.ApiTypeTL) *objectOrPolyTL {
	if len(variants) != 2 {
		return nil
	}
	for i, objectName := range variants {
		polyName := variants[1-i]
		objectScheme, polyScheme := listElements[objectName], listElements[polyName]
		if objectScheme == nil || polyScheme == nil ||
			!utils.IsPolymorphicValue(polyScheme.GetName(), polyScheme.GetDescription()) {
			continue
		}
		for _, objField := range objectScheme.GetFields() {
			if !objField.Optional && len(objField.Types) == 1 && objField.Types[0] == polyName {
				return &objectOrPolyTL{ObjectName: objectName, PolyName: polyName, WrapField: objField.Name}
			}
		}
	}
	return nil
}
