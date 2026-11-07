package component

import "fmt"

func (builder *Context) FillField(fieldName, value string) {
	builder.content += fmt.Sprintf("%s%s: %s,\n", builder.GetTab(), fieldName, value)
}
