package component

import "fmt"

func (builder *Context) AddType(name, typeName string) *Context {
	builder.content += fmt.Sprintf("%stype %s %s", builder.GetTab(), name, typeName)
	return builder
}
