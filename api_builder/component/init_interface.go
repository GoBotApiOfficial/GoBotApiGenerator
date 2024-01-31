package component

import (
	"BotApiCompiler/api_builder/utils"
	"fmt"
)

func (builder *Context) InitInterface(structName string) *Context {
	builder.content += fmt.Sprintf("%stype %s interface {\n", builder.GetTab(), utils.FixStructName(structName))
	builder.tabCount++
	return builder
}
