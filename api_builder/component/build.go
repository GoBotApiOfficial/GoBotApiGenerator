package component

import (
	"BotApiCompiler/consts"
	"fmt"
	"sort"
)

func (builder *Context) Build() []byte {
	var packagesList string
	if len(builder.imports) > 1 {
		packagesList = "import (\n"
		sort.Strings(builder.imports)
		for _, importCode := range builder.imports {
			packagesList += fmt.Sprintf("\t%s\n", importCode)
		}
		packagesList += ")\n\n"
	} else if len(builder.imports) == 1 {
		packagesList = "import " + builder.imports[0] + "\n\n"
	}
	return []byte(fmt.Sprintf(
		"%s\n\npackage %s\n\n%s%s\n%s",
		consts.AGMessage,
		builder.packageName,
		packagesList,
		builder.documentation,
		builder.content,
	))
}
