package utils

import "strings"

func PrettifyField(name string) string {
	nameSplit := strings.Split(name, "_")
	for i, x := range nameSplit {
		nameSplit[i] = strings.ToUpper(x[:1]) + x[1:]
	}
	r := strings.NewReplacer("Id", "ID", "Url", "URL", "Ip", "IP")
	return r.Replace(strings.Join(nameSplit, ""))
}
