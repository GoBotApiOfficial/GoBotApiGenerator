package component

import "fmt"

func (builder *Context) AddFor(condition string) *Context {
	if len(condition) > 0 {
		condition = condition + " "
	}
	builder.content += fmt.Sprintf(
		"%sfor %s{\n",
		builder.GetTab(),
		condition,
	)
	builder.tabCount++
	return builder
}
