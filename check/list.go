package check

import "github.com/jqkard/fn/ds"

func All[T any](items []T, ok func(T) bool) bool {
	for _, item := range items {
		if !ok(item) {
			return false
		}
	}
	return true
}

func Any[T any](items []T, ok func(T) bool) bool {
	for _, item := range items {
		if ok(item) {
			return true
		}
	}
	return false
}

func AllIndex[T any](items []T, ok func(int, T) bool) bool {
	for i, item := range items {
		if !ok(i, item) {
			return false
		}
	}
	return true
}

func AllEqual[T comparable](items []T, value T) bool {
	for _, item := range items {
		if item != value {
			return false
		}
	}
	return true
}

func AnyEqual[T comparable](items []T, value T) bool {
	for _, item := range items {
		if item == value {
			return true
		}
	}
	return false
}

func AllUnique[T comparable](items []T) bool {
	return len(items) == ds.SetFrom(items).Len()
}

func AllSame[T comparable](items []T) bool {
	return ds.SetFrom(items).Len() == 1
}
