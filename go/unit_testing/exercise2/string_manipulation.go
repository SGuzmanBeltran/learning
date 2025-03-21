package exercise2

import (
	"strings"
	"unicode"
)

type StringManipulation struct {}

func (sm *StringManipulation) Reverse(text string) string {
	runes := []rune(text)

	for i, j := 0, len(runes) - 1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
func (sm *StringManipulation) ToUpperCase(text string) string {
	runes := []rune(text)

	for i, r := range runes {
		runes[i] = unicode.ToUpper(r)
	}

	return string(runes)
}
func (sm *StringManipulation) RemoveSpaces(text string) string {
	text = strings.ReplaceAll(text, " ", "")
	text = strings.ReplaceAll(text, "\t", "")
	text = strings.ReplaceAll(text, "\n", "")

	return text
}
