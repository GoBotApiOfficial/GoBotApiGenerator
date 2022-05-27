package component

import (
	"fmt"
)

func (builder *Context) CloseBracket() *Context {
	builder.tabCount--
	builder.content += fmt.Sprintf("%s}\n", builder.GetTab())
	return builder
}
