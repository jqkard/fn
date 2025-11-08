package comb

import (
	"iter"

	"github.com/jqkard/fn"
	"github.com/jqkard/fn/list"
)

func Product[T any](domains ...[]T) iter.Seq2[int, []T] {
	return func(yield func(int, []T) bool) {
		sizes := fn.Map(domains, list.Length)
		total := list.Product(sizes)
		for i := range total {
			tuple := getDomainCombo(domains, sizes, i)
			if !yield(i, tuple) {
				return
			}
		}
	}
}

func ProductRange[T any](start, end int, domains ...[]T) iter.Seq2[int, []T] {
	return func(yield func(int, []T) bool) {
		sizes := fn.Map(domains, list.Length)
		for i := start; i < end; i++ {
			tuple := getDomainCombo(domains, sizes, i)
			if !yield(i, tuple) {
				return
			}
		}
	}
}

func getDomainCombo[T any](domains [][]T, sizes []int, n int) []T {
	numSizes := len(sizes)
	indexes := make([]int, numSizes)
	for i := range numSizes {
		denom := list.Product(sizes[i+1:])
		num := sizes[i] * denom
		indexes[i] = (n % num) / denom
	}
	numDomains := len(domains)
	combo := make([]T, numDomains)
	for i := range numDomains {
		idx := indexes[i]
		combo[i] = domains[i][idx]
	}
	return combo
}
