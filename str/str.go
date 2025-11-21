// The str package contains useful string functions
package str

import (
	"fmt"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var printer = message.NewPrinter(language.English)

func IsEmpty(text string) bool {
	return text == ""
}

func NotEmpty(text string) bool {
	return text != ""
}

func Comma(number int) string {
	return printer.Sprintf("%d", number)
}

func Any[T any](item T) string {
	return fmt.Sprintf("%v", item)
}

func CleanSplit(text string, sep string) []string {
	parts := strings.Split(text, sep)
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}
	return parts
}

func SpaceSplit(text string) []string {
	return strings.Fields(strings.TrimSpace(text))
}

func Lines(text string) []string {
	return CleanSplit(text, "\n")
}

func List[T any](items []T) []string {
	items2 := make([]string, len(items))
	for i, item := range items {
		items2[i] = fmt.Sprintf("%v", item)
	}
	return items2
}

func WrapBraces[T any](items []T) string {
	return "{ " + strings.Join(List(items), ", ") + " }"
}

func WrapBrackets[T any](items []T) string {
	return "[ " + strings.Join(List(items), ", ") + " ]"
}

func WrapParens[T any](items []T) string {
	return "( " + strings.Join(List(items), ", ") + " )"
}
