package component

func (builder *Context) CloseCase() *Context {
	builder.tabCount--
	return builder
}
