package api_builder

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/consts"
	"fmt"
	"os"
	"path"
)

func (ctx *Context) BuildListeners() {
	outputFileFolder := path.Join(consts.OutputFolder, "raw_listeners.go")
	builder := component.NewBuilder()
	builder.SetPackage(utils.MainPackage())
	update := ctx.ApiTL.Types["Update"]
	builder.AddImport("", fmt.Sprintf("%s/types", consts.PackageName))
	builder.AddFunc(
		"ctx *Client",
		"OnRawUpdate",
		[]string{"handler func(update types.Update)"},
		"",
	)
	builder.AddIf("ctx.handlers == nil")
	builder.SetVarValue("ctx.handlers", "make(map[string][]interface{})").AddLine()
	builder.CloseBracket()
	builder.SetVarValue("ctx.handlers[\"raw\"]", "append(ctx.handlers[\"raw\"], handler)").AddLine()
	builder.CloseBracket().AddLine()
	for _, method := range update.GetFields() {
		if method.Name != "update_id" {
			structName := utils.PrettifyField(method.Name)
			genericName := utils.FixGeneric(false, "", method.Types, true, false)
			listenerName := utils.FixName(method.Name)
			builder.AddFunc(
				"ctx *Client",
				fmt.Sprintf("On%s", structName),
				[]string{fmt.Sprintf("handler func(client Client, update %s)", genericName)},
				"",
			)
			builder.AddIf("ctx.handlers == nil")
			builder.SetVarValue("ctx.handlers", "make(map[string][]interface{})").AddLine()
			builder.CloseBracket()
			builder.SetVarValue(
				fmt.Sprintf("ctx.handlers[\"%s\"]", listenerName),
				fmt.Sprintf("append(ctx.handlers[\"%s\"], handler)", listenerName),
			)
			builder.AddLine()
			builder.CloseBracket().AddLine()
		}
	}
	_ = os.WriteFile(outputFileFolder, builder.Build(), 0755)
}
