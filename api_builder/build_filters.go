package api_builder

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/consts"
	"fmt"
	"golang.org/x/exp/slices"
	"path"
	"sort"
	"strings"
)

func (ctx *Context) BuildFilters() {
	type FilterableData struct {
		Name string
		Type string
	}
	var sharedFields []FilterableData
	var sharedTypes []string
	var foundPackages []string
	for _, typeScheme := range ctx.ApiTL.Types["Update"].GetFields() {
		genericName := utils.GenericType(typeScheme.Types, false, false)
		genericAddName := fmt.Sprintf("types.%s", genericName)
		if !utils.IsSimpleGeneric(typeScheme.Types) {
			for _, field := range ctx.ApiTL.Types[genericName].GetFields() {
				if slices.Contains(consts.CommonFields, field.Name) {
					tmpGeneric := fmt.Sprintf("*%s", utils.GenericTypeWithName(field.Types, field.Name, true, false))
					if !slices.Contains(foundPackages, genericAddName) {
						foundPackages = append(foundPackages, genericAddName)
					}
					if !utils.Contains(sharedFields, func(f FilterableData) bool {
						return f.Name == field.Name
					}) {
						sharedFields = append(sharedFields, FilterableData{
							Name: field.Name,
							Type: tmpGeneric,
						})
						sharedTypes = append(sharedTypes, tmpGeneric)
					} else if utils.Contains(sharedFields, func(f FilterableData) bool {
						return f.Name == field.Name && tmpGeneric != f.Type
					}) {
						panic(fmt.Sprintf("Field %s has different types: %s", field.Name, tmpGeneric))
					}
				}
			}
		}
	}
	var sharedFieldsList []string
	var sharedCallable []string
	for _, typeName := range sharedFields {
		sharedFieldsList = append(sharedFieldsList, fmt.Sprintf("%s %s", typeName.Name, typeName.Type))
		sharedCallable = append(sharedCallable, typeName.Name)
	}
	slices.Sort(sharedTypes)
	sort.Slice(sharedFields, func(i, j int) bool {
		return sharedFields[i].Name < sharedFields[j].Name
	})
	slices.Sort(foundPackages)
	slices.Sort(sharedCallable)
	slices.Sort(sharedFieldsList)

	outputFileFolder := path.Join(consts.OutputFolder, "filters", "filterable.go")
	builder := component.NewBuilder()
	builder.SetPackage("filters")
	builder.AddImport("", fmt.Sprintf("%s/types", consts.PackageName))
	builder.InitInterface("Filterable")
	builder.AddInterfaceField(strings.Join(foundPackages, " | \n"))
	builder.CloseBracket()
	utils.WriteCode(outputFileFolder, builder.Build())

	outputFileFolder = path.Join(consts.OutputFolder, "filters", "type.go")
	builder = component.NewBuilder()
	builder.SetPackage("filters")
	builder.AddImport("", fmt.Sprintf("%s", consts.PackageName))
	builder.AddImport("", fmt.Sprintf("%s/types", consts.PackageName))
	builder.InitStruct("DataFilter")
	for _, typeName := range sharedFields {
		builder.AddField(typeName.Name, typeName.Type, "")
	}
	builder.AddField("Client", "*gobotapi.Client", "")
	builder.AddField("RawUpdate", "any", "")
	builder.CloseBracket()
	utils.WriteCode(outputFileFolder, builder.Build())

	outputFileFolder = path.Join(consts.OutputFolder, "filters", "filterable_data.go")
	builder = component.NewBuilder()
	builder.SetPackage("filters")
	builder.AddImport("", fmt.Sprintf("%s", consts.PackageName))
	builder.AddImport("", fmt.Sprintf("%s/types", consts.PackageName))
	builder.AddFunc("", "filterableData", []string{"client *gobotapi.Client", "filterable any"}, fmt.Sprintf("(%s)", "*DataFilter"))
	builder.InitVarValue("dataResult", "&DataFilter{}").AddLine()
	builder.InitSwitch("filterable.(type)")
	for _, typeName := range foundPackages {
		builder.AddCase(false, []string{typeName})
		builder.InitVarValue("x", fmt.Sprintf("filterable.(%s)", typeName)).AddLine()
		pInfo := ctx.ApiTL.Types[strings.TrimPrefix(typeName, "types.")].GetFields()
		for _, field := range pInfo {
			if slices.Contains(consts.CommonFields, field.Name) {
				contentField := fmt.Sprintf("x.%s", utils.PrettifyField(field.Name))
				if !field.Optional {
					contentField = fmt.Sprintf("&%s", contentField)
				}
				builder.SetVarValue(fmt.Sprintf("dataResult.%s", utils.FixStructName(field.Name)), contentField).AddLine()
			}
		}
	}
	builder.CloseBracket()
	builder.SetVarValue("dataResult.Client", "client").AddLine()
	builder.SetVarValue("dataResult.RawUpdate", "filterable").AddLine()
	builder.AddReturn("dataResult")
	builder.CloseBracket()
	utils.WriteCode(outputFileFolder, builder.Build())
}
