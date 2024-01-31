package component

import (
	"strings"
)

func (builder *Context) GetTab() string {
	return strings.Repeat("\t", builder.tabCount)
}
