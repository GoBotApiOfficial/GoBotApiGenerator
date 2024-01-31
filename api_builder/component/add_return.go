package component

import "fmt"

func (builder *Context) AddReturn(returnType string) *Context {
	if len(returnType) > 0 {
		returnType = fmt.Sprintf(" %s", returnType)
	}
	builder.content += fmt.Sprintf("%sreturn%s", builder.GetTab(), returnType)
	return builder
}
