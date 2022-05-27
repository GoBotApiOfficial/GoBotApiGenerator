package api_grabber

import (
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
}
