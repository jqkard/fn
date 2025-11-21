// The lang package contains quality-of-life language functions
package lang

import (
	"strconv"
	"strings"
)

type Number interface {
	~int | ~uint | ~float32 | ~float64
}

func Ternary[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func Abs(x int) int {
	return Ternary(x < 0, -x, x)
}

func ParseInt(value string) int {
	number, err := strconv.Atoi(strings.TrimSpace(value))
	return Ternary(err == nil, number, 0)
}

func ParseFloat(value string) float64 {
	number, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	return Ternary(err == nil, number, 0)
}
