package utils

import "strings"

func GenericTypeWithName(generics []string, varName string, isMethod bool, removeBrackets bool) string {
	r := FixGeneric(false, varName, generics, isMethod, false)
	if removeBrackets {
		r = strings.ReplaceAll(r, "[]", "")
	}
	return r
}

func GenericType(generics []string, isMethod bool, removeBrackets bool) string {
	return GenericTypeWithName(generics, "", isMethod, removeBrackets)
}
