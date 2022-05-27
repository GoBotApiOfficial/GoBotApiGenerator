package api_builder

import (
	"BotApiCompiler/api_builder/utils"
)

func (ctx *Context) Build() *Context {
	utils.InitializeFolders()
	utils.InitializeMod()
	consts := ctx.BuildConstants()
	for _, typeScheme := range ctx.ApiTL.Types {
		BuildComponent(typeScheme, ctx.ApiTL.Types, consts)
	}
	for _, typeScheme := range ctx.ApiTL.Methods {
		BuildComponent(typeScheme, nil, consts)
	}
	ctx.BuildListeners()
	ctx.BuildRun()
	return ctx
}
