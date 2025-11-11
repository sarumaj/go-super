package pythonic

import "errors"

// Set is a Python-like generic set implementation. It requires element type T to be comparable.
// Internally backed by map[T]struct{} for O(1) membership operations.
type Set[T comparable] struct {
	m map[T]struct{}
}

// NewSet creates a new Set with optional initial elements.
func NewSet[T comparable](elems ...T) *Set[T] {
	s := &Set[T]{m: make(map[T]struct{}, len(elems))}
	for _, e := range elems {
		s.m[e] = struct{}{}
	}
	return s
}

// SetFromSlice creates a Set from a slice (duplicates removed).
func SetFromSlice[T comparable](s []T) *Set[T] {
	st := &Set[T]{m: make(map[T]struct{}, len(s))}
	for _, e := range s {
		st.m[e] = struct{}{}
	}
	return st
}

// Len returns number of elements in the set.
func (s *Set[T]) Len() int { return len(s.m) }

// Contains returns true if element is present.
func (s *Set[T]) Contains(v T) bool {
	if s == nil {
		return false
	}
	_, ok := s.m[v]
	return ok
}

// Add inserts element into the set.
func (s *Set[T]) Add(v T) {
	if s.m == nil {
		s.m = make(map[T]struct{})
	}
	s.m[v] = struct{}{}
}

// Update adds multiple elements.
func (s *Set[T]) Update(elems ...T) {
	if s.m == nil {
		s.m = make(map[T]struct{})
	}
	for _, e := range elems {
		s.m[e] = struct{}{}
	}
}

// Remove deletes element and returns error if not present.
func (s *Set[T]) Remove(v T) error {
	if s == nil || s.m == nil {
		return errors.New("value not found")
	}
	if _, ok := s.m[v]; !ok {
		return errors.New("value not found")
	}
	delete(s.m, v)
	return nil
}

// Discard deletes element if present; no error if missing.
func (s *Set[T]) Discard(v T) {
	if s == nil || s.m == nil {
		return
	}
	delete(s.m, v)
}

// Pop removes and returns an arbitrary element. Returns error if set is empty.
func (s *Set[T]) Pop() (T, error) {
	var zero T
	if s == nil || len(s.m) == 0 {
		return zero, errors.New("pop from an empty set")
	}
	for k := range s.m {
		delete(s.m, k)
		return k, nil
	}
	return zero, errors.New("pop from an empty set")
}

// Clear removes all elements.
func (s *Set[T]) Clear() {
	if s == nil {
		return
	}
	s.m = make(map[T]struct{})
}

// AsSlice returns elements as a slice (iteration order is unspecified).
func (s *Set[T]) AsSlice() []T {
	out := make([]T, 0, len(s.m))
	for k := range s.m {
		out = append(out, k)
	}
	return out
}

// Copy returns a shallow copy of the set.
func (s *Set[T]) Copy() *Set[T] {
	if s == nil {
		return &Set[T]{m: make(map[T]struct{})}
	}
	ns := NewSet[T]()
	ns.Update(s.AsSlice()...)
	return ns
}

// Union returns a new set with elements from s and other.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	res := NewSet[T]()
	if s != nil {
		for k := range s.m {
			res.m[k] = struct{}{}
		}
	}
	if other != nil {
		for k := range other.m {
			res.m[k] = struct{}{}
		}
	}
	return res
}

// Intersection returns a new set with elements common to s and other.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	res := NewSet[T]()
	if s == nil || other == nil {
		return res
	}
	// iterate smaller set for efficiency
	var small, large *Set[T]
	if s.Len() <= other.Len() {
		small, large = s, other
	} else {
		small, large = other, s
	}
	for k := range small.m {
		if _, ok := large.m[k]; ok {
			res.m[k] = struct{}{}
		}
	}
	return res
}

// Difference returns a new set with elements in s that are not in other.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	res := NewSet[T]()
	if s == nil {
		return res
	}
	if other == nil {
		return s.Copy()
	}
	for k := range s.m {
		if _, ok := other.m[k]; !ok {
			res.m[k] = struct{}{}
		}
	}
	return res
}

// SymmetricDifference returns elements in either s or other but not both.
func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	a := s.Difference(other)
	b := other.Difference(s)
	return a.Union(b)
}

// IsSubset returns true if s is a subset of other.
func (s *Set[T]) IsSubset(other *Set[T]) bool {
	if s == nil || s.Len() == 0 {
		return true
	}
	if other == nil {
		return false
	}
	for k := range s.m {
		if _, ok := other.m[k]; !ok {
			return false
		}
	}
	return true
}

// IsSuperset returns true if s is a superset of other.
func (s *Set[T]) IsSuperset(other *Set[T]) bool {
	if other == nil || other.Len() == 0 {
		return true
	}
	if s == nil {
		return false
	}
	for k := range other.m {
		if _, ok := s.m[k]; !ok {
			return false
		}
	}
	return true
}

// Equal returns true if s and other have exactly the same elements.
func (s *Set[T]) Equal(other *Set[T]) bool {
	if s == nil && other == nil {
		return true
	}
	if s == nil || other == nil {
		return false
	}
	if s.Len() != other.Len() {
		return false
	}
	for k := range s.m {
		if _, ok := other.m[k]; !ok {
			return false
		}
	}
	return true
}

// UnionWith mutates s to become the union of s and other.
func (s *Set[T]) UnionWith(other *Set[T]) *Set[T] {
	if other == nil {
		return s
	}
	if s.m == nil {
		s.m = make(map[T]struct{})
	}
	for k := range other.m {
		s.m[k] = struct{}{}
	}
	return s
}

// IntersectionWith mutates s to keep only elements present in other.
func (s *Set[T]) IntersectionWith(other *Set[T]) *Set[T] {
	if s == nil || s.m == nil {
		return s
	}
	if other == nil {
		// intersection with empty -> clear
		s.Clear()
		return s
	}
	for k := range s.m {
		if _, ok := other.m[k]; !ok {
			delete(s.m, k)
		}
	}
	return s
}

// DifferenceWith mutates s to remove any elements present in other.
func (s *Set[T]) DifferenceWith(other *Set[T]) *Set[T] {
	if s == nil || s.m == nil || other == nil {
		return s
	}
	for k := range other.m {
		delete(s.m, k)
	}
	return s
}

// SymmetricDifferenceWith mutates s to be the symmetric difference between s and other.
func (s *Set[T]) SymmetricDifferenceWith(other *Set[T]) *Set[T] {
	if other == nil {
		return s
	}
	if s == nil || s.m == nil {
		// become a copy of other
		s.m = make(map[T]struct{}, len(other.m))
		for k := range other.m {
			s.m[k] = struct{}{}
		}
		return s
	}
	for k := range other.m {
		if _, ok := s.m[k]; ok {
			delete(s.m, k)
		} else {
			s.m[k] = struct{}{}
		}
	}
	return s
}

// IsDisjoint returns true if s has no elements in common with other.
func (s *Set[T]) IsDisjoint(other *Set[T]) bool {
	if s == nil || s.Len() == 0 || other == nil || other.Len() == 0 {
		return true
	}
	// iterate smaller set
	if s.Len() <= other.Len() {
		for k := range s.m {
			if _, ok := other.m[k]; ok {
				return false
			}
		}
	} else {
		for k := range other.m {
			if _, ok := s.m[k]; ok {
				return false
			}
		}
	}
	return true
}
