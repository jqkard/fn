package conv

import (
	"strconv"
	"strings"

	"github.com/jqkard/fn"
)

func Abs(x int) int {
	return fn.Ternary(x < 0, -x, x)
}

func ParseInt(value string) int {
	number, err := strconv.Atoi(strings.TrimSpace(value))
	return fn.Ternary(err == nil, number, 0)
}

func ParseFloat(value string) float64 {
	number, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	return fn.Ternary(err == nil, number, 0)
}
