package api_grabber

import (
	"BotApiCompiler/api_grabber/types"
)

func (ctx *Context) ApplySubtypes() {
	var subTypesContainer = make(map[string]string)
	for _, x := range ctx.ApiTL.Types {
		if len(x.Subtypes) > 0 {
			for _, y := range x.Subtypes {
				subTypesContainer[y] = x.Name
			}
		}
	}
	for _, x := range ctx.ApiTL.Types {
		for _, x2 := range ctx.ApiTL.Types {
			if x2.Name == subTypesContainer[x.Name] && x.TypeIds != nil {
				ctx.ApiTL.Types[x2.Name].TypeIds = &types.TypeIdsDescriptor{
					CommonName: x.TypeIds.CommonName,
				}
				break
			}
		}
	}
}
