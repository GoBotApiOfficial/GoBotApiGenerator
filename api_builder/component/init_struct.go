package component

import (
	"BotApiCompiler/api_builder/utils"
	"fmt"
)

func (builder *Context) InitStruct(structName string) *Context {
	builder.content += fmt.Sprintf("%stype %s struct {\n", builder.GetTab(), utils.FixStructName(structName))
	builder.tabCount++
	return builder
}
