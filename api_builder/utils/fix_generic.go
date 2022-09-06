package utils

import (
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
)

func FixGeneric(isOptional bool, varName string, typeName []string, isMethod bool, isRaw bool) string {
	singleTypeName := typeName[0]
	returnGeneric := ""
	if isOptional {
		returnGeneric = "*"
	}
	if len(typeName) > 1 && !(strings.Contains(varName, "id") || slices.Contains(typeName, "InputFile")) {
		arrays, _ := FixArray(singleTypeName)
		interfaceName := "any"
		if varName == "media" && isMethod {
			interfaceName = "types.InputMedia"
		}
		return fmt.Sprintf("%s%s", arrays, interfaceName)
	} else if varName == "chat_id" && slices.Contains(typeName, "Integer") && slices.Contains(typeName, "String") {
		return "any"
	}
	switch singleTypeName {
	case "Integer":
		if strings.Contains(varName, "id") || strings.Contains(varName, "date") {
			return "int64"
		} else {
			return "int"
		}
	case "String":
		return "string"
	case "Boolean":
		return "bool"
	case "Float", "Float number":
		return "float64"
	case "any":
		return "any"
	default:
		if strings.HasPrefix(singleTypeName, "Array of ") {
			arrays, generic := FixArray(singleTypeName)
			return arrays + FixGeneric(false, varName, []string{generic}, isMethod, isRaw)
		}
		if isMethod {
			if isRaw {
				returnGeneric += "rawTypes."
			} else {
				returnGeneric += "types."
			}

		}
		return returnGeneric + singleTypeName
	}
}
