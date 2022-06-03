package utils

func Contains[E comparable](s []E, v func(E) bool) bool {
	for _, vs := range s {
		if v(vs) {
			return true
		}
	}
	return false
}
