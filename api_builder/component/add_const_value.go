package component

import "fmt"

func (builder *Context) AddConstValue(name, value string) {
	if len(value) > 0 {
		value = fmt.Sprintf(" = %s", value)
	}
	builder.content += fmt.Sprintf("%s%s%s\n", builder.GetTab(), name, value)
}
