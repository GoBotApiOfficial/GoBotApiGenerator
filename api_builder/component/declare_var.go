package component

import (
	"fmt"
)

func (builder *Context) DeclareVar(name, typeName string) *Context {
	if typeName == "struct" {
		defer func() {
			builder.tabCount++
		}()
		typeName = fmt.Sprintf("%s {", typeName)
	}
	builder.content += fmt.Sprintf("%svar %s %s", builder.GetTab(), name, typeName)
	return builder
}
