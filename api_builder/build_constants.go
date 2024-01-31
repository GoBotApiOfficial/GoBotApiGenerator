package api_builder

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/consts"
	"golang.org/x/exp/slices"
	"path"
	"sort"
	"strings"
)

func (ctx *Context) BuildConstants() []string {
	outputFileFolder := path.Join(consts.OutputFolder, "types", "kinds.go")
	builder := component.NewBuilder()
	builder.SetPackage("types")
	var existingConstants []string
	for _, typeScheme := range ctx.ApiTL.Types {
		if len(typeScheme.GetSubTypes()) > 0 && !typeScheme.IsSendMethod() {
			builder.AddComment(strings.ToUpper(utils.FixName(typeScheme.GetName())))
			builder.InitConst()
			addedNum := 0
			for _, field := range typeScheme.GetSubTypes() {
				if !slices.Contains(existingConstants, field) {
					existingConstants = append(existingConstants, field)
					if addedNum > 0 {
						builder.AddConstValue("Type"+field, "")
					} else {
						builder.AddConstValue("Type"+field, "iota")
					}
					addedNum++
				}
			}
			builder.CloseRoundBracket()
			builder.AddLine()
		}
	}
	builder.AddComment("RETURN_TYPES")
	builder.InitConst()
	var constTypes []string
	for _, typeScheme := range ctx.ApiTL.Methods {
		for _, field := range typeScheme.GetReturns() {
			if !slices.Contains(constTypes, field) {
				constTypes = append(constTypes, field)
			}
		}
	}
	constTypes = append(constTypes, "ErrorMessage")
	sort.Strings(constTypes)
	addedNum := 0
	for _, field := range constTypes {
		fieldPrettified := utils.FixConstName(field)
		if !slices.Contains(existingConstants, fieldPrettified) {
			if addedNum == 0 {
				builder.AddConstValue("Type"+fieldPrettified, "iota")
			} else {
				builder.AddConstValue("Type"+fieldPrettified, "")
			}
			addedNum++
		}
	}
	builder.CloseRoundBracket()
	utils.WriteCode(outputFileFolder, builder.Build())
	return constTypes
}
