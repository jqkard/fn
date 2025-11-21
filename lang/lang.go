// The lang package contains quality-of-life language functions
package lang

func Ternary[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}
