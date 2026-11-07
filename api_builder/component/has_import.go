package component

import (
	"fmt"
	"strings"
)

func (builder *Context) HasImport(importName string) bool {
	return strings.Contains(fmt.Sprintf("import \"%s\"", builder.Build()), importName)
}
