package sub_build

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/interfaces"
	"BotApiCompiler/api_builder/utils"
	"fmt"
)

func BuildMethodName[Scheme interfaces.SchemeInterface](typeScheme Scheme, builder *component.Context) {
	if typeScheme.GetType() == "methods" {
		builder.AddLine()
		builder.AddFunc(
			utils.FixStructName(typeScheme.GetName()),
			"MethodName",
			nil,
			"string",
		)
		builder.AddReturn(fmt.Sprintf("\"%s\"", typeScheme.GetName())).AddLine()
		builder.CloseBracket()
	}
}
