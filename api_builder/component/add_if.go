package component

import "fmt"

func (builder *Context) AddIf(condition string) *Context {
	builder.content += fmt.Sprintf("%sif %s {\n", builder.GetTab(), condition)
	builder.tabCount++
	return builder
}
