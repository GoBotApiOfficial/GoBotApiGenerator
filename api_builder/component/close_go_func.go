package component

import (
	"fmt"
	"strings"
)

func (builder *Context) CloseGoFunc(funcParams []string) *Context {
	builder.tabCount--
	builder.content += fmt.Sprintf("%s}(%s)\n", builder.GetTab(), strings.Join(funcParams, ", "))
	return builder
}
