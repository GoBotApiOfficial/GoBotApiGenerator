package component

func (builder *Context) InitConst() {
	builder.content += "const (\n"
	builder.tabCount++
}
