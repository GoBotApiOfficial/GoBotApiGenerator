package component

import "fmt"

func (builder *Context) AddComment(comment string) {
	builder.content += fmt.Sprintf("// %s\n", comment)
}
