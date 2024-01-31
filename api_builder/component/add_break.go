package component

import "fmt"

func (builder *Context) AddBreak() *Context {
	builder.content += fmt.Sprintf("%sbreak\n", builder.GetTab())
	return builder
}
