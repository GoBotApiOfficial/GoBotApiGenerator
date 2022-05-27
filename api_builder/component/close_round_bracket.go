package component

func (builder *Context) CloseRoundBracket() {
	builder.tabCount--
	builder.content += builder.GetTab() + ")\n"
}
