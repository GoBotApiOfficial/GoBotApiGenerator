package api_grabber

import (
	"BotApiCompiler/api_grabber/types"
	"BotApiCompiler/consts"
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
	"strings"
)

func (ctx *Context) DownloadApiTL() *types.ApiTL {
	resp, err := soup.Get(consts.BotApiLink)
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	devRules := doc.Find("div", "id", "dev_page_content")
	var currName string
	var isMethod bool
	var version string
	for _, x := range devRules.Children() {
		if x.NodeValue == "p" && len(version) == 0 {
			fullText := x.FullText()
			if strings.HasPrefix(fullText, "Bot API") {
				version = strings.TrimSpace(strings.TrimPrefix(fullText, "Bot API"))
			}
		}
		if x.NodeValue == "h3" || x.NodeValue == "hr" {
			currName = ""
			isMethod = false
		}
		if x.NodeValue == "h4" {
			anchor := x.Find("a")
			name := anchor.Attrs()["name"]
			if strings.Contains(name, "-") {
				currName = ""
				isMethod = false
				continue
			}
			currName, isMethod = ctx.GetNameAndType(x)
		}
		if len(currName) == 0 {
			continue
		}
		if x.NodeValue == "p" {
			desc := x.FullText()
			if isMethod && len(ctx.ApiTL.Methods[currName].Description) == 0 {
				ctx.GetReturns(currName, desc)
			}
			if isMethod {
				ctx.ApiTL.Methods[currName].Description = append(ctx.ApiTL.Methods[currName].Description, desc)
			} else {
				ctx.ApiTL.Types[currName].Description = append(ctx.ApiTL.Types[currName].Description, desc)
			}
		}
		if x.NodeValue == "table" {
			ctx.GetFields(currName, isMethod, x)
		}
		if x.NodeValue == "ul" {
			ctx.GetSubtypes(currName, x)
		}
	}
	ctx.ApplySubtypes()
	ctx.CheckSenderStruct()
	for _, x := range ctx.ApiTL.Methods {
		if x.Returns == nil {
			panic(fmt.Sprintf("No return found for %s", x.Name))
		}
	}
	ctx.ApiTL.Version = version
	return ctx.ApiTL
}
