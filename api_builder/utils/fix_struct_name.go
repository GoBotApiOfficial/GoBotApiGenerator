package utils

import "strings"

func FixStructName(name string) string {
	return strings.ToUpper(name[:1]) + name[1:]
}
