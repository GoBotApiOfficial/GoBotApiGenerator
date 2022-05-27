package utils

import "strings"

func FixArray(generic string) (string, string) {
	countArrayOf := strings.Count(generic, "Array of ")
	cleanedGeneric := strings.Replace(generic, "Array of ", "", countArrayOf)
	return strings.Repeat("[]", countArrayOf), cleanedGeneric
}
