package component

import (
	"fmt"
)

func (builder *Context) InitSwitch(switchValue string) *Context {
	builder.content += fmt.Sprintf(
		"%sswitch %s {\n",
		builder.GetTab(),
		switchValue,
	)
	builder.tabCount++
	return builder
}
