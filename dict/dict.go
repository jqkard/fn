// The dict package contains useful map functions
package dict

import (
	"maps"
	"slices"
)

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func Keys[K comparable, V any](items map[K]V) []K {
	return slices.Collect(maps.Keys(items))
}

func Values[K comparable, V any](items map[K]V) []V {
	return slices.Collect(maps.Values(items))
}

func Entries[K comparable, V any](items map[K]V) []Entry[K, V] {
	entries := make([]Entry[K, V], 0, len(items))
	for k, v := range items {
		entries = append(entries, Entry[K, V]{k, v})
	}
	return entries
}

func HasKey[K comparable, V any](items map[K]V, key K) bool {
	_, ok := items[key]
	return ok
}

func NoKey[K comparable, V any](items map[K]V, key K) bool {
	return !HasKey(items, key)
}

func SetDefault[K comparable, V any](items map[K]V, key K, defaultValue V) {
	if _, ok := items[key]; !ok {
		items[key] = defaultValue
	}
}

func GetOr[K comparable, V any](items map[K]V, key K, defaultValue V) V {
	if value, ok := items[key]; ok {
		return value
	}
	return defaultValue
}

func TallyValues[K, V comparable](items map[K]V, values []V) map[V]int {
	count := make(map[V]int, len(values))
	for _, value := range values {
		count[value] = 0
	}
	for _, value := range items {
		if NoKey(count, value) {
			continue
		}
		count[value] += 1
	}
	return count
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
