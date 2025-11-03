package list

func IntRange(start, end int) []int {
	numbers := make([]int, end-start)
	for n := start; n < end; n++ {
		numbers[n-start] = n
	}
	return numbers
}
