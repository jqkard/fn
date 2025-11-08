package list

import "github.com/jqkard/fn"

func IntRange(start, end int) []int {
	numbers := make([]int, end-start)
	for n := start; n < end; n++ {
		numbers[n-start] = n
	}
	return numbers
}

func Sum[T fn.Number](items []T) T {
	var total T = 0
	for _, item := range items {
		total += item
	}
	return total
}

func Product[T fn.Number](items []T) T {
	var product T = 1
	for _, item := range items {
		product *= item
	}
	return product
}

func Length[T any](items []T) int {
	return len(items)
}

func IsEmpty[T any](items []T) bool {
	return len(items) == 0
}

func NotEmpty[T any](items []T) bool {
	return len(items) > 0
}
