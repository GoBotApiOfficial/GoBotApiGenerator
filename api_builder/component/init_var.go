package component

import (
	"fmt"
)

func (builder *Context) InitVarValue(varName, varValue string) *Context {
	builder.content += fmt.Sprintf("%s%s := %s", builder.GetTab(), varName, varValue)
	return builder
}
