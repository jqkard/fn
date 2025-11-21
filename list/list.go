// The list package contains useful list functions
package list

import (
	"math/rand/v2"

	"github.com/jqkard/fn/number"
)

func IntRange(start, end int) []int {
	numbers := make([]int, end-start)
	for n := start; n < end; n++ {
		numbers[n-start] = n
	}
	return numbers
}

func Sum[T number.Number](items []T) T {
	var total T = 0
	for _, item := range items {
		total += item
	}
	return total
}

func Product[T number.Number](items []T) T {
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

func TallyItems[T comparable](items []T) map[T]int {
	count := make(map[T]int)
	for _, item := range items {
		count[item] += 1
	}
	return count
}

func Count[T comparable](items []T, value T) int {
	count := 0
	for _, item := range items {
		if item == value {
			count += 1
		}
	}
	return count
}
