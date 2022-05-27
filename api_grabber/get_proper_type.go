package api_grabber

func GetProperType(t string) string {
	switch t {
	case "Messages":
		return "Message"
	case "Float numbers":
		return "Float"
	case "Int":
		return "Integer"
	case "True", "Bool":
		return "Boolean"
	default:
		return t
	}
}
