package utils

func IsSimpleGeneric(generics []string) bool {
	generic := GenericType(generics, false, false)
	switch generic {
	case "int", "int32", "int64", "float32", "float64", "string", "bool", "interface{}":
		return true
	}
	return false
}
