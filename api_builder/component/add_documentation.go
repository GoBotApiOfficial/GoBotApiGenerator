package component

import (
	"strings"
)

func (builder *Context) AddDocumentation(name string, documentation []string) *Context {
	documentation[0] = name + " " + documentation[0]
	var description []string
	for _, sentence := range documentation {
		if strings.Contains(sentence, ".") {
			for _, subSentence := range strings.Split(sentence, ". ") {
				if len(subSentence) > 0 {
					description = append(description, "// "+strings.Replace(subSentence, "This object represents", "Represents", 1))
				}
			}
		} else {
			description = append(description, "// "+strings.Replace(sentence, "This object represents", "Represents", 1))
		}
	}
	builder.documentation += strings.Join(description, "\n")
	return builder
}
