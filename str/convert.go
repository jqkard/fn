package str

import "fmt"

func Any[T any](item T) string {
	return fmt.Sprintf("%v", item)
}
