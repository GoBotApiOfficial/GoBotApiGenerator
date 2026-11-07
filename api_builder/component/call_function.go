package component

import (
	"fmt"
	"strings"
)

func (builder *Context) CallFunction(funcName string, args []string) *Context {
	builder.content += fmt.Sprintf("%s%s(%s)", builder.GetTab(), funcName, strings.Join(args, ", "))
	return builder
}
