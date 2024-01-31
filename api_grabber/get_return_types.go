package api_grabber

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func GetReturnTypes(descPart string) []string {
	arrayMatch, _ := regexp.Compile("(?i)(?:array of )+(\\w*)")
	matches := arrayMatch.FindStringSubmatch(descPart)
	if len(matches) > 0 {
		ret := CleanTgType(matches[1])
		var rets []string
		for _, x := range ret {
			rets = append(rets, fmt.Sprintf("Array of %s", x))
		}
		return rets
	} else {
		words := strings.Split(descPart, " ")
		var rets []string
		for _, ret := range words {
			for _, r := range ret {
				if unicode.IsUpper(r) {
					rets = append(rets, CleanTgType(ret)...)
					break
				}
			}
		}
		return rets
	}
}
