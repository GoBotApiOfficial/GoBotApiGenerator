package component

import "fmt"

func (builder *Context) CloseOpenBracket() {
	builder.tabCount--
	builder.content += fmt.Sprintf("%s} {\n", builder.GetTab())
	builder.tabCount++
}
