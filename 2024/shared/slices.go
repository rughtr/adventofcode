package shared

import "iter"

// Map applies a mapping function to each element of a slice and returns a sequence of transformed elements.
func Map[Slice ~[]E, E any, T any](s Slice, mapFunc func(value E) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range s {
			mv := mapFunc(v)
			if !yield(mv) {
				return
			}
		}
	}
}

// Unique returns a sequence of unique elements from the provided slice, maintaining their original order.
func Unique[Slice ~[]E, E comparable](s Slice) iter.Seq[E] {
	return func(yield func(E) bool) {
		seen := make(map[E]struct{})
		for _, v := range s {
			_, found := seen[v]
			if found {
				continue
			}
			seen[v] = struct{}{}
			if !yield(v) {
				return
			}
		}
	}
}
