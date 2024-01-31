package component

import (
	"BotApiCompiler/api_builder/utils"
	"fmt"
)

func (builder *Context) AddField(name, genericName, jsonName string) *Context {
	jsonTmp := ""
	if len(jsonName) > 0 {
		jsonTmp = fmt.Sprintf(" `json:\"%s\"`", jsonName)
	}
	builder.content += fmt.Sprintf(
		"%s%s %s%s\n",
		builder.GetTab(),
		utils.PrettifyField(name),
		genericName,
		jsonTmp,
	)
	return builder
}
