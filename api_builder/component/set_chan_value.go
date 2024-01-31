package component

import "fmt"

func (builder *Context) SetChanValue(name string, value string) *Context {
	builder.content += fmt.Sprintf("%s%s <- %s", builder.GetTab(), name, value)
	return builder
}
