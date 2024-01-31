package component

import (
	"fmt"
	"strings"
)

func (builder *Context) AddFunc(structName, funcName string, funcParams []string, funcReturn string) *Context {
	structOption := ""
	if len(structName) > 0 {
		structOption = fmt.Sprintf("(%s) ", structName)
	}
	returnOption := ""
	if len(funcReturn) > 0 {
		returnOption = fmt.Sprintf(" %s", funcReturn)
	}
	builder.content += fmt.Sprintf(
		"%sfunc %s%s(%s)%s {\n",
		builder.GetTab(),
		structOption,
		funcName,
		strings.Join(funcParams, ", "),
		returnOption,
	)
	builder.tabCount++
	return builder
}
