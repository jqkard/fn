package dict

import (
	"maps"
	"slices"
)

func Keys[K comparable, V any](items map[K]V) []K {
	return slices.Collect(maps.Keys(items))
}

func Values[K comparable, V any](items map[K]V) []V {
	return slices.Collect(maps.Values(items))
}

func HasKey[K comparable, V any](items map[K]V, key K) bool {
	_, ok := items[key]
	return ok
}

func NoKey[K comparable, V any](items map[K]V, key K) bool {
	return !HasKey(items, key)
}

func TallyValues[K, V comparable](items map[K]V, values []V) map[V]int {
	tally := make(map[V]int, len(values))
	for _, value := range values {
		tally[value] = 0
	}
	for _, value := range items {
		if NoKey(tally, value) {
			continue
		}
		tally[value] += 1
	}
	return tally
}

func Zip[K comparable, V any](keys []K, values []V) map[K]V {
	m := make(map[K]V, len(keys))
	numValues := len(values)
	for i, k := range keys {
		if i >= numValues {
			break
		}
		m[k] = values[i]
	}
	return m
}

func Counter[T comparable](items []T) map[T]int {
	count := make(map[T]int)
	for _, item := range items {
		count[item] = 0
	}
	return count
}

func Flags[T comparable](items []T, flag bool) map[T]bool {
	flags := make(map[T]bool)
	for _, item := range items {
		flags[item] = flag
	}
	return flags
}
