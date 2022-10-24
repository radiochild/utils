package text

import (
	"strings"
)

func AppendText(text ...string) string {
	var sb strings.Builder
	for _, textItem := range text {
		trimmedText := strings.TrimSpace(textItem)
		if len(trimmedText) > 0 {
			sb.WriteString(trimmedText)
			sb.WriteString(" ")
		}
	}
	return strings.TrimSpace(sb.String())
}
