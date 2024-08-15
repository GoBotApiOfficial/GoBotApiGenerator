package sub_build

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/api_grabber/types"
	"BotApiCompiler/consts"
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
)

func BuildCheck(builder *component.Context, isMethod bool, sendChildTypes map[string]types.FieldTL) {
	for fieldName, fieldTypes := range sendChildTypes {
		genericNameTmp := utils.FixGeneric(false, fieldTypes.Name, fieldTypes.Types, false, false)
		arrayCounts := strings.Count(genericNameTmp, "[]")
		entityName := fmt.Sprintf("entity.%s", utils.PrettifyField(fieldName))
		for i := 0; i < arrayCounts; i++ {
			builder.AddFor(fmt.Sprintf("_, x%d := range %s", i, entityName))
			entityName = fmt.Sprintf("x%d", i)
		}
		originalEntityName := entityName
		isChatIDWithUsername := fieldTypes.Name == "chat_id" && slices.Contains(fieldTypes.Types, "Integer") && slices.Contains(fieldTypes.Types, "String")
		if fieldTypes.Optional || isChatIDWithUsername {
			builder.AddImport("", "reflect")
			builder.AddIf(fmt.Sprintf("!reflect.ValueOf(%s).IsNil()", entityName))
			if !strings.Contains(genericNameTmp, "any") && len(fieldTypes.FullTypes) == 0 {
				entityName = fmt.Sprintf("(*%s)", entityName)
			}
		}
		builder.InitSwitch(fmt.Sprintf("%s.(type)", entityName))
		var fixedCases []string
		var listTypes []string
		if len(fieldTypes.FullTypes) > 0 {
			listTypes = fieldTypes.FullTypes
		} else {
			listTypes = fieldTypes.Types
		}
		for _, field := range listTypes {
			genericName := utils.GenericType([]string{field}, false, true)
			if isMethod {
				builder.AddImport("", fmt.Sprintf("%s/types", consts.PackageName))
				if utils.IsSimpleGeneric([]string{genericName}) {
					fixedCases = append(fixedCases, genericName)
					if genericName == "int" {
						fixedCases = append(fixedCases, "int64")
					}
				} else {
					fixedCases = append(fixedCases, fmt.Sprintf("*types.%s", genericName))
				}
			} else {
				fixedCases = append(fixedCases, genericName)
			}
		}
		builder.AddCase(false, fixedCases)
		builder.AddBreak().CloseCase()
		builder.AddDefault()
		builder.AddImport("", "fmt")
		builder.AddReturn(fmt.Sprintf("nil, fmt.Errorf(\"%s: unknown type: %%T\", %s)", fieldName, originalEntityName))
		builder.CloseCase().AddLine()
		for i := 0; i < arrayCounts; i++ {
			builder.CloseBracket()
		}
		if fieldTypes.Optional || isChatIDWithUsername {
			builder.CloseBracket()
		}
		builder.CloseBracket()
	}
}
