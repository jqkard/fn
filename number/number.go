// The number package contains useful number functions
package number

type Number interface {
	~int | ~uint | ~float32 | ~float64
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
