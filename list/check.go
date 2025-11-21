package list

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

func AnyIndex[T any](items []T, ok func(int, T) bool) bool {
	for i, item := range items {
		if ok(i, item) {
			return true
		}
	}
	return false
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

func AllTrue(items []bool) bool {
	return AllEqual(items, true)
}

func AllFalse(items []bool) bool {
	return AllEqual(items, false)
}

func AnyTrue(items []bool) bool {
	return AnyEqual(items, true)
}

func AnyFalse(items []bool) bool {
	return AnyEqual(items, false)
}

func AllSame[T comparable](items []T) bool {
	return len(TallyItems(items)) == 1
}

func AllUnique[T comparable](items []T) bool {
	return len(TallyItems(items)) == len(items)
}
