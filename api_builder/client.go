package api_builder

import (
	"BotApiCompiler/api_grabber/types"
)

func Client(apiScheme *types.ApiTL) *Context {
	return &Context{apiScheme}
}
