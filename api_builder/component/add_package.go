package component

import "BotApiCompiler/api_builder/utils"

func (builder *Context) SetPackage(packageName string) *Context {
	builder.packageName = utils.FixName(packageName)
	return builder
}
