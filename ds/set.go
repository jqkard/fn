package ds

import "strings"

type Set[T comparable] struct {
	items map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]bool)}
}

func SetFrom[T comparable](items []T) *Set[T] {
	set := NewSet[T]()
	set.AddItems(items)
	return set
}

func (s *Set[T]) Add(item T) {
	s.items[item] = true
}

func (s *Set[T]) AddItems(items []T) {
	for _, item := range items {
		s.Add(item)
	}
}

func (s *Set[T]) Delete(item T) {
	delete(s.items, item)
}

func (s Set[T]) Contains(item T) bool {
	return s.items[item]
}

func (s Set[T]) Len() int {
	return len(s.items)
}

func (s Set[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s Set[T]) Items() []T {
	items := make([]T, 0, len(s.items))
	for item := range s.items {
		items = append(items, item)
	}
	return items
}

func (s1 Set[T]) Union(s2 *Set[T]) *Set[T] {
	s3 := NewSet[T]()
	for x := range s1.items {
		s3.Add(x)
	}
	for x := range s2.items {
		s3.Add(x)
	}
	return s3
}

func (s1 Set[T]) Intersection(s2 *Set[T]) *Set[T] {
	s3 := NewSet[T]()
	for x := range s1.items {
		if s2.Contains(x) {
			s3.Add(x)
		}
	}
	return s3
}

func (s1 Set[T]) Difference(s2 *Set[T]) *Set[T] {
	s3 := NewSet[T]()
	for x := range s1.items {
		if !s2.Contains(x) {
			s3.Add(x)
		}
	}
	return s3
}

type Subsets struct {
	Universal []string
	Names     []string
	Subsets   [][]string
}

func NewSubsets(universal string, subsetLines []string) *Subsets {
	numSubsets := len(subsetLines)
	names := make([]string, numSubsets)
	subsets := make([][]string, numSubsets)
	for i, line := range subsetLines {
		parts := strings.Split(line, ":")
		names[i] = strings.TrimSpace(parts[0])
		subsets[i] = strings.Fields(strings.TrimSpace(parts[1]))
	}
	return &Subsets{
		Universal: strings.Fields(strings.TrimSpace(universal)),
		Names:     names,
		Subsets:   subsets,
	}
}
