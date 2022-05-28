package api_grabber

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"strings"
)

func (ctx *Context) GetSubtypes(currName string, x soup.Root) {
	var subtypes []string
	var isSend bool
	for _, li := range x.FindAll("li") {
		SubtypeName := li.FullText()
		if strings.HasPrefix(SubtypeName, "Input") {
			isSend = true
		}
		subtypes = append(subtypes, SubtypeName)
	}
	ctx.ApiTL.Types[currName].IsSend = isSend
	ctx.ApiTL.Types[currName].Subtypes = subtypes
	for _, subtype := range subtypes {
		ctx.ApiTL.Types[currName].Description = append(ctx.ApiTL.Types[currName].Description, fmt.Sprintf(" - %s", subtype))
	}
}
