package component

import (
	"fmt"
	"strings"
)

func (builder *Context) AddCase(isString bool, caseValue []string) *Context {
	var caseValueTmp []string
	for _, value := range caseValue {
		if isString {
			caseValueTmp = append(caseValueTmp, fmt.Sprintf("\"%s\"", value))
		} else {
			caseValueTmp = append(caseValueTmp, value)
		}
	}
	builder.content += fmt.Sprintf(
		"%scase %s:\n",
		builder.GetTab(),
		strings.Join(caseValueTmp, ", "),
	)
	builder.tabCount++
	return builder
}
