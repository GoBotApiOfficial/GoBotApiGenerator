package component

import "fmt"

func (builder *Context) AddInterfaceField(name string) *Context {
	builder.content += fmt.Sprintf("%s%s\n", builder.GetTab(), name)
	return builder
}
