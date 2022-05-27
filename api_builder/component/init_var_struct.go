package component

import "fmt"

func (builder *Context) InitVarStruct(name string, structName string) {
	builder.InitVarValue(name, fmt.Sprintf("%s {", structName)).AddLine()
	builder.tabCount++
}
