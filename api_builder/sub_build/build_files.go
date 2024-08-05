package sub_build

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/interfaces"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/api_grabber/types"
	"BotApiCompiler/consts"
	"fmt"
	"regexp"
	"strings"
)

func BuildFiles[Scheme interfaces.SchemeInterface](typeScheme Scheme, builder *component.Context, filesInput []types.FieldTL) {
	isMethod := typeScheme.GetType() == "methods"
	isMediaInput := consts.MediaInputRgx.MatchString(typeScheme.GetName())
	mediaFieldRgx := regexp.MustCompile("^\\[].*\\.Input.*Media$")
	structName := utils.FixStructName(typeScheme.GetName())
	structName = "*" + structName
	if consts.GenericInputRgx.MatchString(typeScheme.GetName()) {
		builder.InitInterface(typeScheme.GetName())
		builder.AddInterfaceFunc("SetAttachment", []string{"string"}, "")
		builder.AddInterfaceFunc("SetAttachmentThumb", []string{"string"}, "")
		builder.CloseBracket()
	} else if (len(filesInput) > 0 && isMethod) || isMediaInput {
		builder.AddLine()
		if isMethod {
			builder.AddFunc(
				fmt.Sprintf("entity %s", structName),
				"ProgressCallable",
				nil,
				"rawTypes.ProgressCallable",
			)
			builder.AddReturn("entity.Progress").AddLine()
			builder.CloseBracket()
		}
		builder.AddLine()
		builder.AddFunc(
			fmt.Sprintf("entity %s", structName),
			"Files",
			nil,
			"map[string]rawTypes.InputFile",
		)
		builder.InitVarValue("files", "make(map[string]rawTypes.InputFile)").AddLine()
		builder.AddImport("rawTypes", fmt.Sprintf("%s/types/raw", consts.PackageName))
		for _, field := range filesInput {
			genericName := utils.FixGeneric(field.Optional, field.Name, field.Types, isMethod, true)
			prettifiedField := utils.PrettifyField(field.Name)
			typeFieldName := strings.ToLower(consts.InputTypeNameRgx.FindAllStringSubmatch(typeScheme.GetName(), -1)[0][2])

			if strings.Contains(genericName, "InputFile") {
				builder.InitSwitch(fmt.Sprintf("entity.%s.(type)", prettifiedField))
				if isMethod {
					builder.AddImport("", fmt.Sprintf("%s/types", consts.PackageName))
					builder.AddCase(false, []string{"types.InputBytes"})
				} else {
					builder.AddCase(false, []string{"InputBytes"})
				}
				if field.Name == "media" {
					builder.SetVarValue(fmt.Sprintf("files[\"%s\"]", typeFieldName), "entity.Media").AddLine()
				} else {
					builder.SetVarValue(fmt.Sprintf("files[\"%s\"]", field.Name), fmt.Sprintf("entity.%s", prettifiedField)).AddLine()
				}
				if isMethod {
					if field.Name == "thumbnail" {
						builder.AddImport("", fmt.Sprintf("%s/types", consts.PackageName))
						builder.SetVarValue(fmt.Sprintf("entity.%s", prettifiedField), "types.InputURL(\"attach://thumbnail\")").AddLine()
					} else {
						builder.SetVarValue(fmt.Sprintf("entity.%s", prettifiedField), "nil").AddLine()
					}
				}
				builder.CloseCase().CloseBracket()
			} else if genericName == "string" {
				typeFieldName = strings.ToLower(consts.InputTypeNameRgx.FindAllStringSubmatch(typeScheme.GetName(), -1)[0][2])
				builder.AddIf(fmt.Sprintf("entity.%s == \"\"", prettifiedField))
				builder.SetVarValue(fmt.Sprintf("files[\"%s\"]", typeFieldName), "entity.Media").AddLine()
				builder.CloseBracket()
			} else if mediaFieldRgx.MatchString(genericName) {
				builder.AddFor(fmt.Sprintf("i, x0 := range entity.%s", prettifiedField))
				builder.InitVarValue("x1", "x0.(rawTypes.InputMediaFiles).Files()").AddLine()
				builder.AddFor("k, v := range x1")
				builder.DeclareVar("attachName", "string").AddLine()
				builder.AddIf("k == \"thumbnail\"")
				builder.SetVarValue("attachName", "fmt.Sprintf(\"file-%d-thumbnail\", i)").AddLine()
				builder.CallFunction("x0.SetAttachmentThumb", []string{"attachName"}).AddLine()
				builder.AddElse()
				builder.SetVarValue("attachName", "fmt.Sprintf(\"file-%d\", i)").AddLine()
				builder.CallFunction("x0.SetAttachment", []string{"attachName"}).AddLine()
				builder.CloseBracket()
				builder.SetVarValue("files[attachName]", "v").AddLine()
				builder.CloseBracket().CloseBracket()
			} else {
				tmpField := prettifiedField
				if field.Name == "media" {
					tmpField = "Media.(rawTypes.InputMediaFiles)"
				}
				builder.AddFor(fmt.Sprintf("k, v := range entity.%s.Files()", tmpField))
				builder.SetVarValue("files[k]", "v").AddLine()
				builder.AddIf("k == \"thumbnail\"")
				builder.CallFunction("entity.Media.SetAttachmentThumb", []string{"k"}).AddLine()
				builder.AddElse()
				builder.CallFunction("entity.Media.SetAttachment", []string{"k"}).AddLine()
				builder.CloseBracket()
				builder.CloseBracket()
			}
		}
		builder.AddReturn("files").AddLine()
		builder.CloseBracket()
		if isMediaInput {
			builder.AddImport("", "fmt")
			builder.AddLine()
			builder.AddFunc("entity "+structName, "SetAttachment", []string{"attach string"}, "")
			builder.SetVarValue("entity.Media", "InputURL(fmt.Sprintf(\"attach://%s\", attach))").AddLine()
			builder.CloseBracket().AddLine()
			var foundThumb bool
			for _, field := range typeScheme.GetFields() {
				if field.Name == "thumbnail" {
					foundThumb = true
					break
				}
			}
			if foundThumb {
				builder.AddFunc("entity "+structName, "SetAttachmentThumb", []string{"attach string"}, "")
				builder.SetVarValue("entity.Thumbnail", "InputURL(fmt.Sprintf(\"attach://%s\", attach))").AddLine()
			} else {
				builder.AddFunc("entity "+structName, "SetAttachmentThumb", []string{"_ string"}, "")
			}
			builder.CloseBracket()
		}
	} else if isMethod {
		builder.AddImport("rawTypes", fmt.Sprintf("%s/types/raw", consts.PackageName))
		builder.AddLine()
		builder.AddFunc(
			fmt.Sprintf("entity %s", structName),
			"ProgressCallable",
			nil,
			"rawTypes.ProgressCallable",
		)
		builder.AddReturn("nil").AddLine()
		builder.CloseBracket()
		builder.AddLine()
		builder.AddFunc("entity "+structName, "Files", nil, "map[string]rawTypes.InputFile")
		builder.AddReturn("map[string]rawTypes.InputFile{}").AddLine()
		builder.CloseBracket()
	}
}
