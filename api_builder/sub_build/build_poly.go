package sub_build

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/utils"
	"fmt"
)

type AliasFieldTL struct {
	Name      string
	Generic   string
	JsonName  string
	ParseFunc string
}

func BuildPolyValue(builder *component.Context, name string) {
	builder.AddImport("", "encoding/json")
	builder.AddImport("", "bytes")
	builder.AddLine().AddLine()
	builder.AddComment(fmt.Sprintf("%sValue is either %sPlain, %sArray or %s", name, name, name, name))
	builder.InitInterface(fmt.Sprintf("%sValue", name))
	builder.AddInterfaceFunc("Kind", nil, "int")
	builder.CloseBracket().AddLine()
	builder.AddType(fmt.Sprintf("%sPlain", name), "string").AddLine().AddLine()
	builder.AddFunc(fmt.Sprintf("x %sPlain", name), "Kind", nil, "int")
	builder.AddReturn(fmt.Sprintf("Type%sPlain", name)).AddLine()
	builder.CloseBracket().AddLine()
	builder.AddType(fmt.Sprintf("%sArray", name), fmt.Sprintf("[]%sValue", name)).AddLine().AddLine()
	builder.AddFunc(fmt.Sprintf("x %sArray", name), "Kind", nil, "int")
	builder.AddReturn(fmt.Sprintf("Type%sArray", name)).AddLine()
	builder.CloseBracket().AddLine()
	builder.AddFunc("", fmt.Sprintf("Parse%sValue", name), []string{"data []byte"}, fmt.Sprintf("(%sValue, error)", name))
	builder.SetVarValue("data", "bytes.TrimSpace(data)").AddLine()
	builder.AddIf("len(data) == 0 || string(data) == \"null\"")
	builder.AddReturn("nil, nil").AddLine()
	builder.CloseBracket()
	builder.AddIf("data[0] == '\"'")
	builder.DeclareVar("plain", "string").AddLine()
	builder.AddIf("err := json.Unmarshal(data, &plain); err != nil")
	builder.AddReturn("nil, err").AddLine()
	builder.CloseBracket()
	builder.AddReturn(fmt.Sprintf("%sPlain(plain), nil", name)).AddLine()
	builder.CloseBracket()
	builder.AddIf("data[0] == '['")
	builder.DeclareVar("rawItems", "[]json.RawMessage").AddLine()
	builder.AddIf("err := json.Unmarshal(data, &rawItems); err != nil")
	builder.AddReturn("nil, err").AddLine()
	builder.CloseBracket()
	builder.InitVarValue("items", fmt.Sprintf("make(%sArray, 0, len(rawItems))", name)).AddLine()
	builder.AddFor("_, rawItem := range rawItems")
	builder.InitVarValue("item, err", fmt.Sprintf("Parse%sValue(rawItem)", name)).AddLine()
	builder.AddIf("err != nil")
	builder.AddReturn("nil, err").AddLine()
	builder.CloseBracket()
	builder.SetVarValue("items", "append(items, item)").AddLine()
	builder.CloseBracket()
	builder.AddReturn("items, nil").AddLine()
	builder.CloseBracket()
	builder.DeclareVar("item", name).AddLine()
	builder.AddIf("err := json.Unmarshal(data, &item); err != nil")
	builder.AddReturn("nil, err").AddLine()
	builder.CloseBracket()
	builder.AddReturn("item, nil").AddLine()
	builder.CloseBracket()
}

func BuildPolyUnmarshal(builder *component.Context, structName string, fields []AliasFieldTL) {
	builder.AddImport("", "encoding/json")
	builder.AddLine()
	builder.AddFunc(fmt.Sprintf("entity *%s", structName), "UnmarshalJSON", []string{"data []byte"}, "error")
	builder.DeclareVar("alias", "struct").AddLine()
	for _, field := range fields {
		genericName := field.Generic
		if len(field.ParseFunc) > 0 {
			genericName = "json.RawMessage"
		}
		builder.AddField(field.Name, genericName, field.JsonName)
	}
	builder.CloseBracket()
	builder.AddIf("err := json.Unmarshal(data, &alias); err != nil")
	builder.AddReturn("err").AddLine()
	builder.CloseBracket()
	builder.DeclareVar("err", "error").AddLine()
	for _, field := range fields {
		prettyName := utils.PrettifyField(field.Name)
		if len(field.ParseFunc) > 0 {
			builder.AddIf(fmt.Sprintf("entity.%s, err = %s(alias.%s); err != nil", prettyName, field.ParseFunc, prettyName))
			builder.AddReturn("err").AddLine()
			builder.CloseBracket()
		} else {
			builder.SetVarValue(fmt.Sprintf("entity.%s", prettyName), fmt.Sprintf("alias.%s", prettyName)).AddLine()
		}
	}
	builder.AddReturn("nil").AddLine()
	builder.CloseBracket()
}
