package component

import "fmt"

func (builder *Context) AddDefault() *Context {
	builder.content += fmt.Sprintf("%sdefault:\n", builder.GetTab())
	builder.tabCount++
	return builder
}
