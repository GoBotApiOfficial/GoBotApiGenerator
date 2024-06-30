package api_grabber

import (
	"fmt"
	"regexp"
	"strings"
)

func (ctx *Context) GetReturns(methodName, description string) {
	retSearch, _ := regexp.Compile("(?i)(?:on success,|returns)([^.]*)(?:on success)?")
	matches := retSearch.FindAllStringSubmatch(description, -1)
	var retTypes []string
	if len(matches) > 0 {
		size := len(matches) - 1
		sizeSub := len(matches[size]) - 1
		retTypes = GetReturnTypes(strings.TrimSpace(matches[size][sizeSub]))
	} else {
		panic(fmt.Sprintf("No return found for %s", methodName))
	}
	ctx.ApiTL.Methods[methodName].Returns = retTypes
}
