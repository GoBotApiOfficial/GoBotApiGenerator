package api_grabber

import (
	"BotApiCompiler/api_grabber/types"
	"github.com/anaskhan96/soup"
	"unicode"
)

func (ctx *Context) GetNameAndType(x soup.Root) (string, bool) {
	currName := x.Text()
	isMethod := !unicode.IsUpper(rune(x.Text()[0]))
	if isMethod {
		ctx.ApiTL.Methods[currName] = &types.ApiMethodTL{
			Name: currName,
		}
	} else {
		ctx.ApiTL.Types[currName] = &types.ApiTypeTL{
			Name: currName,
		}
	}
	return currName, isMethod
}
