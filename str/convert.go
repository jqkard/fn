package str

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var printer = message.NewPrinter(language.English)

func Comma(number int) string {
	return printer.Sprintf("%d", number)
}

func Any[T any](item T) string {
	return fmt.Sprintf("%v", item)
}
