package api_grabber

import "BotApiCompiler/api_grabber/types"

func Client() *Context {
	return &Context{
		ApiTL: &types.ApiTL{
			Methods: make(map[string]*types.ApiMethodTL),
			Types:   make(map[string]*types.ApiTypeTL),
		},
	}
}
