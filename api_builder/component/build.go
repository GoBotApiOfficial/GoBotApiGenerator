package component

import (
	"fmt"
	"strings"
)

func (builder *Context) Build() []byte {
	var packagesList string
	if len(builder.imports) > 0 {
		packagesList = strings.Join(builder.imports, "\n")
		packagesList += "\n\n"
	}
	return []byte(fmt.Sprintf("package %s\n\n%s%s\n%s", builder.packageName, packagesList, builder.documentation, builder.content))
}
