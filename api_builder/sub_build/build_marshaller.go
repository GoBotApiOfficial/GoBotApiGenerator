package sub_build

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/interfaces"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/api_grabber/types"
	"BotApiCompiler/consts"
	"fmt"
	"strings"
)

func BuildMarshaller[Scheme interfaces.SchemeInterface](typeScheme Scheme, builder *component.Context, listElements map[string]*types.ApiTypeTL) {
	sendChildTypes := make(map[string]types.FieldTL)
	isMethod := typeScheme.GetType() == "methods"
	for _, field := range typeScheme.GetFields() {
		genericName := strings.ReplaceAll(utils.FixGeneric(false, field.Name, field.Types, false, false), "[]", "")
		if listElements[genericName] != nil {
			genericCheck := listElements[genericName]
			if genericCheck.IsSendMethod() && len(genericCheck.GetFields()) == 0 && genericName != "InputFile" {
				sendChildTypes[field.Name] = field
			}
		} else if genericName == "interface{}" {
			sendChildTypes[field.Name] = field
		}
	}
	if strings.HasPrefix(typeScheme.GetName(), "InputMedia") {
		builder.AddLine()
		parentStructName := utils.FixStructName(typeScheme.GetName())
		inputField := strings.ToLower(strings.TrimPrefix(parentStructName, "InputMedia"))
		builder.AddImport("", "encoding/json")
		builder.AddFunc(
			fmt.Sprintf("entity %s", parentStructName),
			"MarshalJSON",
			nil,
			"([]byte, error)",
		)
		builder.InitInlineStruct("alias")
		for _, field := range typeScheme.GetFields() {
		GenerateGeneric:
			isFieldRawImport := utils.IsRawField(field.Types)
			if field.Types[0] == "InputFile" {
				field.Optional = false
			}
			genericName := utils.FixGeneric(
				field.Optional,
				field.Name,
				field.Types,
				isMethod || isFieldRawImport,
				isFieldRawImport,
			)
			jsonName := field.Name
			if field.Optional || field.Types[0] == "InputFile" {
				jsonName += ",omitempty"
			}
			if field.Name == "media" && genericName == "string" {
				field = types.FieldTL{
					Name: field.Name,
					Types: []string{
						"InputFile",
					},
					Optional: field.Optional,
				}
				goto GenerateGeneric
			}
			builder.AddField(
				field.Name,
				genericName,
				jsonName,
			)
		}
		builder.CloseOpenBracket()
		for _, field := range typeScheme.GetFields() {
			prettifiedField := utils.PrettifyField(field.Name)
			var value string
			if field.Name == "type" {
				value = fmt.Sprintf("\"%s\"", inputField)
			} else {
				value = fmt.Sprintf("entity.%s", prettifiedField)
			}
			builder.FillField(
				prettifiedField,
				value,
			)
		}
		builder.CloseBracket()
		builder.AddReturn("json.Marshal(alias)").AddLine()
		builder.CloseBracket()
	} else if len(sendChildTypes) > 0 {
		builder.AddLine()
		parentStructName := utils.FixStructName(typeScheme.GetName())
		builder.AddImport("", "encoding/json")
		builder.AddImport("", "fmt")
		builder.AddFunc(
			fmt.Sprintf("entity %s", parentStructName),
			"MarshalJSON",
			nil,
			"([]byte, error)",
		)
		for fieldName, fieldTypes := range sendChildTypes {
			genericNameTmp := utils.FixGeneric(false, fieldTypes.Name, fieldTypes.Types, false, false)
			arrayCounts := strings.Count(genericNameTmp, "[]")
			entityName := fmt.Sprintf("entity.%s", utils.PrettifyField(fieldName))
			for i := 0; i < arrayCounts; i++ {
				builder.AddFor(fmt.Sprintf("_, x%d := range %s", i, entityName))
				entityName = fmt.Sprintf("x%d", i)
			}
			originalEntityName := entityName
			if fieldTypes.Optional {
				builder.AddIf(fmt.Sprintf("%s != nil", entityName))
				if !strings.Contains(genericNameTmp, "interface") {
					entityName = fmt.Sprintf("(*%s)", entityName)
				}
			}
			builder.InitSwitch(fmt.Sprintf("%s.(type)", entityName))
			var fixedCases []string
			for _, field := range fieldTypes.Types {
				genericName := strings.ReplaceAll(utils.FixGeneric(false, "", []string{field}, false, false), "[]", "")
				if isMethod {
					builder.AddImport("", fmt.Sprintf("%s/types", consts.PackageName))
					fixedCases = append(fixedCases, fmt.Sprintf("*types.%s", genericName))
				} else {
					fixedCases = append(fixedCases, genericName)
				}
			}
			builder.AddCase(false, fixedCases)
			builder.AddBreak().CloseCase()
			builder.AddDefault()
			builder.AddReturn(fmt.Sprintf("nil, fmt.Errorf(\"%s: unknown type: %%T\", %s)", fieldName, originalEntityName))
			builder.CloseCase().AddLine()
			for i := 0; i < arrayCounts; i++ {
				builder.CloseBracket()
			}
			if fieldTypes.Optional {
				builder.CloseBracket()
			}
			builder.CloseBracket()
		}
		builder.AddType("x0", parentStructName).AddLine()
		builder.AddReturn("json.Marshal((x0)(entity))").AddLine()
		builder.CloseBracket()
	}
}
