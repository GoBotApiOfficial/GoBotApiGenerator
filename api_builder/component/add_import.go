package component

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func (builder *Context) AddImport(alias, packageName string) *Context {
	tmpAlias := ""
	if len(alias) > 0 {
		tmpAlias = fmt.Sprintf("%s ", alias)
	}
	importCode := fmt.Sprintf("%s\"%s\"", tmpAlias, packageName)
	if !slices.Contains(builder.imports, importCode) {
		builder.imports = append(builder.imports, importCode)
	}
	return builder
}
