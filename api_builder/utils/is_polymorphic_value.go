package utils

import (
	"fmt"
	"strings"
)

func IsPolymorphicValue(name string, description []string) bool {
	for _, line := range description {
		if strings.Contains(line, "either a String") &&
			strings.Contains(line, fmt.Sprintf("Array of %s", name)) {
			return true
		}
	}
	return false
}
