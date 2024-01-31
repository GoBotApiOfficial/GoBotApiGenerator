package api_grabber

import (
	"strings"
)

func CleanTgType(t string) []string {
	var pref string
	if strings.HasPrefix(t, "Array of ") {
		pref = "Array of "
		t = t[len("Array of "):]
	}
	var fixedOrs []string
	for _, or := range strings.Split(t, " or ") {
		fixedOrs = append(fixedOrs, strings.TrimSpace(or))
	}
	var fixedAnds []string
	for _, and := range strings.Split(strings.Join(fixedOrs, " and "), " and ") {
		fixedAnds = append(fixedAnds, strings.TrimSpace(and))
	}
	var fixedComma []string
	for _, comma := range strings.Split(strings.Join(fixedAnds, ", "), ", ") {
		fixedComma = append(fixedComma, strings.TrimSpace(comma))
	}
	var cleanedTgType []string
	for _, x := range fixedComma {
		cleanedTgType = append(cleanedTgType, pref+GetProperType(x))
	}
	return cleanedTgType
}
