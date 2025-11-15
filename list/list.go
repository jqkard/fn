package list

import (
	"math/rand"

	"github.com/jqkard/fn"
)

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

func Shuffle[T any](items []T) {
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})
}

func Divide(numItems, numParts int) [][2]int {
	portion := numItems / numParts
	ranges := make([][2]int, numParts)
	for i := 0; i < numParts-1; i++ {
		start := i * portion
		ranges[i] = [2]int{start, start + portion}
	}
	i := numParts - 1
	start := i * portion
	ranges[i] = [2]int{start, numItems}
	return ranges
}

func CountValue[T comparable](items []T, value T) int {
	count := 0
	for _, item := range items {
		if item == value {
			count += 1
		}
	}
	return count
}
