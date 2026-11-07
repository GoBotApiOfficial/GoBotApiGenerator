package component

import "fmt"

func (builder *Context) InitInlineStruct(varName string) {
	builder.content += fmt.Sprintf("\t%s := struct {\n", varName)
	builder.tabCount++
}
