package component

import (
	"fmt"
	"strings"
)

func (builder *Context) AddInterfaceFunc(funcName string, funcParams []string, funcReturn string) *Context {
	builder.content += fmt.Sprintf(
		"%s%s (%s) %s\n",
		builder.GetTab(),
		funcName,
		strings.Join(funcParams, ", "),
		funcReturn,
	)
	return builder
}
