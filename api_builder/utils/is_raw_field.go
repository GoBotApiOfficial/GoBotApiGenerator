package utils

func IsRawField(generics []string) bool {
	return GenericType(generics, false, false) == "InputFile"
}
