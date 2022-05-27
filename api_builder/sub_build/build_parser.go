package sub_build

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/interfaces"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/consts"
	"fmt"
	"golang.org/x/exp/slices"
)

func BuildParser[Scheme interfaces.SchemeInterface](typeScheme Scheme, builder *component.Context, listConsts []string) {
	if typeScheme.GetType() == "methods" {
		structName := utils.FixStructName(typeScheme.GetName())
		returns := typeScheme.GetReturns()
		var constName string
		if len(returns) > 1 {
			constName = utils.Remove(returns, "Boolean")[0]
		} else {
			constName = returns[0]
		}
		returnName := utils.FixGeneric(false, "", []string{constName}, true, false)
		constName = utils.FixConstName(constName)
		if len(returns) > 1 && !slices.Contains(returns, "Boolean") {
			panic(fmt.Sprintf("Method %s returns more than one value", typeScheme.GetName()))
		} else {
			builder.AddLine()
			builder.AddImport("rawTypes", fmt.Sprintf("%s/types/raw", consts.PackageName))
			builder.AddImport("", fmt.Sprintf("%s/types", consts.PackageName))
			builder.AddImport("", "encoding/json")
			builder.AddFunc(
				structName,
				"ParseResult",
				[]string{"response []byte"},
				"(*rawTypes.Result, error)",
			)
			if len(returns) > 1 {
				builder.DeclareVar("x0", "struct").AddLine()
				builder.AddField("Result", "bool", "result")
				builder.CloseBracket()
				builder.SetVarValue("_", "json.Unmarshal(response, &x0)").AddLine()
				builder.AddIf("x0.Result")
				builder.InitVarStruct("result", "rawTypes.Result")
				builder.FillField("Kind", "types.TypeBoolean")
				builder.FillField("Result", "true")
				builder.CloseBracket()
				builder.AddReturn("&result, nil").AddLine()
				builder.AddElse()
			}
			builder.DeclareVar("x1", "struct").AddLine()
			builder.AddField("Result", returnName, "result")
			builder.CloseBracket()
			builder.InitVarValue("err", "json.Unmarshal(response, &x1)").AddLine()
			builder.AddIf("err != nil")
			builder.AddReturn("nil, err").AddLine()
			builder.CloseBracket()
			builder.InitVarStruct("result", "rawTypes.Result")
			builder.FillField("Kind", fmt.Sprintf("types.Type%s", constName))
			builder.FillField("Result", "x1.Result")
			builder.CloseBracket()
			builder.AddReturn("&result, nil").AddLine()
			if len(returns) > 1 {
				builder.CloseBracket()
			}
			builder.CloseBracket()
		}
	}
}
