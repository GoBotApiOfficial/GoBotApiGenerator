package component

import "fmt"

func (builder *Context) AddLog(typeName, log string) *Context {
	builder.content += fmt.Sprintf("%slog.%s(%s)", builder.GetTab(), typeName, log)
	return builder
}
