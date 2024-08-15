package component

import (
	"fmt"
	"strings"
)

func (builder *Context) InitVarValue(varName, varValue string) *Context {
	builder.content += fmt.Sprintf("%s%s := %s", builder.GetTab(), varName, varValue)
	if strings.HasPrefix(varValue, "func") {
		builder.tabCount++
	}
	return builder
}
