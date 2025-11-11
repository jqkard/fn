package ds

import (
	"strings"

	"github.com/jqkard/fn/dict"
	"github.com/jqkard/fn/str"
)

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
	return dict.Keys(s.items)
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
		parts := str.CleanSplit(line, ":")
		names[i] = parts[0]
		subsets[i] = strings.Fields(parts[1])
	}
	return &Subsets{
		Universal: strings.Fields(universal),
		Names:     names,
		Subsets:   subsets,
	}
}
