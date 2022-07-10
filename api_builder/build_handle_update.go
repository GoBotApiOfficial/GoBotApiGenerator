package api_builder

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/consts"
	"fmt"
	"path"
)

func (ctx *Context) BuildHandleUpdate() {
	outputFileFolder := path.Join(consts.OutputFolder, "handle_update.go")
	update := ctx.ApiTL.Types["Update"]
	builder := component.NewBuilder()
	builder.SetPackage(utils.MainPackage())
	builder.AddImport("", fmt.Sprintf("%s/types", consts.PackageName))
	builder.AddFunc("ctx *BasicClient", "handleUpdate", []string{"user *types.User", "token string", "update types.Update"}, "")
	builder.InitVarStruct("client", "&Client")
	builder.FillField("Token", "token")
	builder.FillField("BasicClient", "ctx")
	builder.FillField("me", "user")
	builder.CloseBracket()
	builder.AddFor("_, x0 := range ctx.handlers[\"raw\"]")
	builder.CallFunction("ctx.concurrencyManager.Wait", nil).AddLine()
	builder.InitGoFunc([]string{"x any"})
	builder.CallFunction("x.(func(*Client, types.Update))", []string{"client", "update"}).AddLine()
	builder.CallFunction("ctx.concurrencyManager.Done", nil).AddLine()
	builder.CloseGoFunc([]string{"x0"})
	builder.CloseBracket()
	for _, method := range update.GetFields() {
		if method.Name != "update_id" {
			structName := utils.PrettifyField(method.Name)
			genericName := utils.GenericType(method.Types, true, false)
			builder.AddIf(fmt.Sprintf("update.%s != nil", structName))
			builder.AddFor(fmt.Sprintf("_, x0 := range ctx.handlers[\"%s\"]", utils.FixName(method.Name)))
			builder.CallFunction("ctx.concurrencyManager.Wait", nil).AddLine()
			builder.InitGoFunc([]string{"x any"})
			builder.CallFunction(fmt.Sprintf("x.(func(*Client, %s))", genericName), []string{"client", fmt.Sprintf("*update.%s", structName)}).AddLine()
			builder.CallFunction("ctx.concurrencyManager.Done", nil).AddLine()
			builder.CloseGoFunc([]string{"x0"})
			builder.CloseBracket().CloseBracket()
		}
	}
	builder.CloseBracket()
	utils.WriteCode(outputFileFolder, builder.Build())
}
