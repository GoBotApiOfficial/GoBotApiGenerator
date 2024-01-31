package component

func (builder *Context) WithContext(n *Context) *Context {
	builder.tabCount = n.tabCount
	return builder
}
