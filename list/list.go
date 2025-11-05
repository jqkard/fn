package list

type number interface {
	~int | ~uint | ~float32 | ~float64
}

func IntRange(start, end int) []int {
	numbers := make([]int, end-start)
	for n := start; n < end; n++ {
		numbers[n-start] = n
	}
	return numbers
}

func Sum[T number](items []T) T {
	var total T = 0
	for _, item := range items {
		total += item
	}
	return total
}

func Product[T number](items []T) T {
	var product T = 1
	for _, item := range items {
		product *= item
	}
	return product
}

func Length[T any](items []T) int {
	return len(items)
}
