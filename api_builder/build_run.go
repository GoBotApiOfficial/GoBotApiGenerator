package api_builder

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/consts"
	"fmt"
	"os"
	"path"
)

func (ctx *Context) BuildRun() {
	outputFileFolder := path.Join(consts.OutputFolder, "run.go")
	update := ctx.ApiTL.Types["Update"]
	builder := component.NewBuilder()
	builder.SetPackage(utils.MainPackage())
	builder.AddImport("", fmt.Sprintf("%s/methods", consts.PackageName))
	builder.AddImport("", fmt.Sprintf("%s/types", consts.PackageName))
	builder.AddImport("", "log")
	builder.AddImport("", "time")
	builder.AddImport("", "net/http")
	builder.AddFunc("ctx *Client", "Run", nil, "")
	builder.AddIf("ctx.isStarted")
	builder.AddReturn("").AddLine()
	builder.CloseBracket()
	builder.SetVarValue("ctx.client", "&http.Client{}").AddLine()
	builder.SetVarValue("ctx.apiURL", "ctx.BotApiConfig.link()").AddLine()
	builder.AddIf("ctx.PollingTimeout == 0")
	builder.SetVarValue("ctx.PollingTimeout", "time.Second * 15").AddLine()
	builder.CloseBracket()
	builder.SetVarValue("ctx.isStarted", "true").AddLine()
	builder.AddIf("ctx.waitStart != nil")
	builder.SetChanValue("ctx.waitStart", "true").AddLine()
	builder.CloseBracket()
	builder.InitVarValue("res, err", "ctx.Invoke(&methods.GetMe{})").AddLine()
	builder.AddIf("err != nil")
	builder.AddLog("Fatal", "err").AddLine()
	builder.CloseBracket()
	builder.SetVarValue("ctx.botID", "res.Result.(types.User).ID").AddLine()
	builder.SetVarValue("ctx.botUsername", "res.Result.(types.User).Username").AddLine()
	builder.CallFunction("showNotice", nil).AddLine()
	builder.AddIf("ctx.NoUpdates")
	builder.AddReturn("").AddLine()
	builder.CloseBracket()
	builder.AddFor("")
	builder.InitVarStruct("getUpdates", "&methods.GetUpdates")
	builder.FillField("Timeout", "int(ctx.PollingTimeout.Seconds())")
	builder.FillField("Offset", "ctx.lastUpdateID")
	builder.CloseBracket()
	builder.InitVarValue("rawUpdates, err", "ctx.Invoke(getUpdates)").AddLine()
	builder.AddIf("!ctx.isStarted")
	builder.AddBreak()
	builder.CloseBracket()
	builder.AddIf("err != nil")
	builder.AddLog("Printf", "\"[%d] Retrying \\\"getUpdates\\\" due to Telegram says %s\", ctx.botID, err").AddLine()
	builder.CallFunction("time.Sleep", []string{"time.Second * 5"}).AddLine()
	builder.AddContinue()
	builder.CloseBracket()
	builder.InitVarValue("updates", "rawUpdates.Result.([]types.Update)").AddLine()
	builder.AddFor("_, update := range updates")
	builder.SetVarValue("ctx.lastUpdateID", "int(update.UpdateID) + 1").AddLine()
	builder.AddFor("_, x0 := range ctx.handlers[\"raw\"]")
	builder.CallFunction("go x0.(func(Client, types.Update))", []string{"*ctx", "update"}).AddLine()
	builder.CloseBracket()
	for _, method := range update.GetFields() {
		if method.Name != "update_id" {
			structName := utils.PrettifyField(method.Name)
			genericName := utils.FixGeneric(false, "", method.Types, true, false)
			builder.AddIf(fmt.Sprintf("update.%s != nil", structName))
			builder.AddFor(fmt.Sprintf("_, x0 := range ctx.handlers[\"%s\"]", utils.FixName(method.Name)))
			builder.CallFunction(fmt.Sprintf("go x0.(func(Client, %s))", genericName), []string{"*ctx", fmt.Sprintf("*update.%s", structName)}).AddLine()
			builder.CloseBracket().CloseBracket()
		}
	}
	builder.CloseBracket().CloseBracket().CloseBracket()
	_ = os.WriteFile(outputFileFolder, builder.Build(), 0755)
}
