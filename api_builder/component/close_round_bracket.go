package component

func (builder *Context) CloseRoundBracket() *Context {
	builder.tabCount--
	builder.content += builder.GetTab() + ")\n"
	return builder
}
