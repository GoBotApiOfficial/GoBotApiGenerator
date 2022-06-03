package component

func (builder *Context) InitConst() *Context {
	builder.content += "const (\n"
	builder.tabCount++
	return builder
}
