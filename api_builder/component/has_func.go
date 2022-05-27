package component

import (
	"fmt"
	"strings"
)

func (builder *Context) HasFunc(funcName string) bool {
	return strings.Contains(fmt.Sprintf("func %s", builder.Build()), funcName)
}
