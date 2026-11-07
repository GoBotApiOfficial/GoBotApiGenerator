package component

func (builder *Context) HaveBody() bool {
	return len(builder.content) > 0
}
