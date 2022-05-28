package component

import "fmt"

func (builder *Context) AddContinue() *Context {
	builder.content += fmt.Sprintf("%scontinue\n", builder.GetTab())
	return builder
}
