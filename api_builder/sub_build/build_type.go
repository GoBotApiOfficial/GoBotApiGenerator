package sub_build

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/interfaces"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/api_grabber/types"
	"BotApiCompiler/consts"
	"fmt"
	"sort"
	"strings"
)

func BuildType[Scheme interfaces.SchemeInterface](typeScheme Scheme, builder *component.Context, listElements map[string]*types.ApiTypeTL) []types.FieldTL {
	var filesInput []types.FieldTL
	var structName string
	isMethod := typeScheme.GetType() == "methods"
	structName = utils.FixStructName(typeScheme.GetName())
	if len(typeScheme.GetFields()) == 0 {
		if !isMethod {
			if !consts.GenericInputRgx.MatchString(typeScheme.GetName()) {
				builder.AddType(utils.FixStructName(structName), "any").AddLine()
			}
		} else {
			builder.AddType(utils.FixStructName(structName), "struct{}").AddLine()
		}
	} else {
		typesOrdered := make([]types.FieldTL, 0, len(typeScheme.GetFields()))
		for _, field := range typeScheme.GetFields() {
			typesOrdered = append(typesOrdered, field)
		}
		sort.Slice(typesOrdered, func(i, j int) bool {
			return typesOrdered[i].Name < typesOrdered[j].Name
		})
		builder.InitStruct(structName)
		for _, field := range typesOrdered {
			jsonName := field.Name
			if field.Types[0] == "InputFile" || consts.GenericInputRgx.MatchString(field.Types[0]) {
				field.Optional = false
			}
			typeName := utils.GenericType(field.Types, false, false)
			if listElements[typeName] != nil && !(len(listElements[typeName].GetSubTypes()) > 0 && !listElements[typeName].IsSendMethod()) {
				if len(listElements[typeName].GetSubTypes()) > 0 {
					field = types.FieldTL{
						Name:     field.Name,
						Types:    []string{"any"},
						Optional: field.Optional,
						Default:  field.Default,
					}
				}
			}
		CheckGeneric:
			isFieldRawImport := utils.IsRawField(field.Types)
			if isFieldRawImport {
				builder.AddImport("rawTypes", fmt.Sprintf("%s/types/raw", consts.PackageName))
			} else if !utils.IsSimpleGeneric(field.Types) && isMethod {
				builder.AddImport("", fmt.Sprintf("%s/types", consts.PackageName))
			}
			genericName := utils.FixGeneric(
				field.Optional,
				field.Name,
				field.Types,
				isMethod || isFieldRawImport,
				isFieldRawImport,
			)
			if len(field.Default) > 0 {
				if isMethod {
					panic("Default value for methods is not supported")
				}
				continue
			}
			if field.Optional || field.Types[0] == "InputFile" || strings.Contains(field.Types[0], "Array of") {
				jsonName += ",omitempty"
			}
			if field.Types[0] == "InputFile" {
				filesInput = append(filesInput, field)
			} else if field.Name == "media" || field.Name == "thumbnail" {
				if genericName == "string" {
					field = types.FieldTL{
						Name: field.Name,
						Types: []string{
							"InputFile",
						},
					}
					goto CheckGeneric
				}
				filesInput = append(filesInput, field)
			}
			builder.AddField(
				field.Name,
				genericName,
				jsonName,
			)
		}
		if len(filesInput) > 0 && isMethod {
			builder.AddImport("rawTypes", fmt.Sprintf("%s/types/raw", consts.PackageName))
			builder.AddField("Progress", "rawTypes.ProgressCallable", "-")
		}
		builder.CloseBracket()
	}
	return filesInput
}
