package utils

import "strings"

func GenericType(generics []string, isMethod bool, removeBrackets bool) string {
	r := FixGeneric(false, "", generics, isMethod, false)
	if removeBrackets {
		r = strings.ReplaceAll(r, "[]", "")
	}
	return r
}
