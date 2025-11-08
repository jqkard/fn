package fn

type Number interface {
	~int | ~uint | ~float32 | ~float64
}

func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func Map[T any, V any](items []T, convert func(T) V) []V {
	results := make([]V, len(items))
	for i, item := range items {
		results[i] = convert(item)
	}
	return results
}

func MapIndex[T any, V any](items []T, convert func(int, T) V) []V {
	results := make([]V, len(items))
	for i, item := range items {
		results[i] = convert(i, item)
	}
	return results
}

func MapList[T any](indexes []int, values []T) []T {
	numValues := len(values)
	results := make([]T, len(indexes))
	for i, idx := range indexes {
		if 0 <= idx && idx < numValues {
			results[i] = values[idx]
		}
	}
	return results
}

func Translate[K comparable, V any](items []K, mask map[K]V) []V {
	results := make([]V, len(items))
	for i, item := range items {
		results[i] = mask[item]
	}
	return results
}

func Filter[T any](items []T, keep func(T) bool) []T {
	results := make([]T, 0, len(items))
	for _, item := range items {
		if keep(item) {
			results = append(results, item)
		}
	}
	return results
}

func FilterIndex[T any](items []T, keep func(int, T) bool) []T {
	results := make([]T, 0, len(items))
	for i, item := range items {
		if keep(i, item) {
			results = append(results, item)
		}
	}
	return results
}
