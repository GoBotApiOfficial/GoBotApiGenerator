package component

import (
	"fmt"
	"strings"
)

func (builder *Context) InitGoFunc(funcParams []string) *Context {
	builder.content += fmt.Sprintf("%sgo func (%s) {\n", builder.GetTab(), strings.Join(funcParams, ", "))
	builder.tabCount++
	return builder
}
