package utils

import "strings"

func FixConstName(name string) string {
	tmpName := strings.Split(name, " ")
	for i, part := range tmpName {
		tmpName[i] = strings.ToUpper(part[:1]) + part[1:]
	}
	return strings.Join(tmpName, "")
}
