package utils

import (
	"strings"
)

func IsRawField(generics []string) bool {
	generic := strings.ReplaceAll(FixGeneric(false, "", generics, false, false), "[]", "")
	return generic == "InputFile"
}
