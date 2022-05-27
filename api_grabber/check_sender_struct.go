package api_grabber

import (
	"strings"
)

func (ctx *Context) CheckSenderStruct() {
	for _, typeScheme := range ctx.ApiTL.Methods {
		for _, method := range typeScheme.GetFields() {
			field := strings.ReplaceAll(method.Types[0], "Array of ", "")
			if ctx.ApiTL.Types[field] != nil {
				ctx.ApiTL.Types[field].IsSend = true
			}
		}
	}
}
