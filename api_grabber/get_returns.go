package api_grabber

import (
	"fmt"
	"regexp"
	"strings"
)

func (ctx *Context) GetReturns(methodName, description string) {
	retSearch, _ := regexp.Compile("(?i)(?:on success,|returns)([^.]*)(?:on success)?")
	retSearch2, _ := regexp.Compile("(?i)([^.]*)(is returned)")
	matches := retSearch.FindStringSubmatch(description)
	matches2 := retSearch2.FindStringSubmatch(description)
	var retTypes []string
	if len(matches) > 0 {
		retTypes = GetReturnTypes(strings.TrimSpace(matches[1]))
	} else if len(matches2) > 0 {
		retTypes = GetReturnTypes(strings.TrimSpace(matches2[1]))
	} else {
		panic(fmt.Sprintf("No return found for %s", methodName))
	}
	ctx.ApiTL.Methods[methodName].Returns = retTypes
}
