package component

func (builder *Context) AddLine() *Context {
	builder.content += "\n"
	return builder
}
