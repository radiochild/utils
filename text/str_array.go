package text

import (
	"strings"
)

func AllToChar(strs []string, c rune) []string {
	updatedStrs := []string{}
	repStr := string(c)
	for _, str := range strs {
		updatedStrs = append(updatedStrs, strings.Repeat(repStr, len(str)))
	}
	return updatedStrs
}

func TabString(strs []string) string {
	return strings.Join(strs, "\t")
}
