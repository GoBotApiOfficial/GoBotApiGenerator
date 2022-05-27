package component

func (builder *Context) AddElse() *Context {
	builder.tabCount--
	builder.content += builder.GetTab() + "} else {\n"
	builder.tabCount++
	return builder
}
