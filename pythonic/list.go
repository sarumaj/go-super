package pythonic

import (
	"errors"
	"reflect"
	"sort"
)

// List is a Python-like generic list implementation.
// It wraps a slice and provides common list methods (append, extend, insert, remove, pop, etc.).
type List[T any] struct {
	data []T
}

// NewList creates a new List with optional initial elements.
func NewList[T any](elems ...T) *List[T] {
	d := make([]T, len(elems))
	copy(d, elems)
	return &List[T]{data: d}
}

// ListFromSlice wraps an existing slice (makes a copy) into a List.
func ListFromSlice[T any](s []T) *List[T] {
	d := make([]T, len(s))
	copy(d, s)
	return &List[T]{data: d}
}

// Len returns the number of elements.
func (l *List[T]) Len() int { return len(l.data) }

// AsSlice returns a copy of the underlying slice.
func (l *List[T]) AsSlice() []T {
	out := make([]T, len(l.data))
	copy(out, l.data)
	return out
}

// helper: normalize index (supports negative indices like Python)
func (l *List[T]) normIndex(i int) int {
	n := len(l.data)
	if i < 0 {
		i = n + i
	}
	return i
}

func (l *List[T]) Concat(other any) *List[T] {
	if other == nil {
		return l.Copy()
	}
	switch o := other.(type) {
	case *List[T]:
		newData := make([]T, len(l.data)+len(o.data))
		copy(newData, l.data)
		copy(newData[len(l.data):], o.data)
		return &List[T]{data: newData}

	case []T:
		newData := make([]T, len(l.data)+len(o))
		copy(newData, l.data)
		copy(newData[len(l.data):], o)
		return &List[T]{data: newData}

	default:
		// unsupported type: return a copy of the original list
		return l.Copy()

	}
}

// Get returns element at index i. Supports negative indices. Returns error if out of range.
func (l *List[T]) Get(i int) (T, error) {
	var zero T
	idx := l.normIndex(i)
	if idx < 0 || idx >= len(l.data) {
		return zero, errors.New("index out of range")
	}
	return l.data[idx], nil
}

// Set assigns value at index i. Supports negative indices.
func (l *List[T]) Set(i int, v T) error {
	idx := l.normIndex(i)
	if idx < 0 || idx >= len(l.data) {
		return errors.New("index out of range")
	}
	l.data[idx] = v
	return nil
}

// Append adds an element to the end.
func (l *List[T]) Append(v T) { l.data = append(l.data, v) }

// Extend appends multiple elements.
func (l *List[T]) Extend(elems ...T) { l.data = append(l.data, elems...) }

// Insert inserts element at index i (before current element at i). Supports negative indices.
func (l *List[T]) Insert(i int, v T) error {
	n := len(l.data)
	if i < 0 {
		i = n + i
	}
	if i < 0 {
		i = 0
	}
	if i > n {
		i = n
	}
	l.data = append(l.data[:i], append([]T{v}, l.data[i:]...)...)
	return nil
}

// Index returns the index of the first occurrence of v. Returns error if not found.
func (l *List[T]) Index(v T) (int, error) {
	for i, it := range l.data {
		if reflect.DeepEqual(it, v) {
			return i, nil
		}
	}
	return -1, errors.New("value not found")
}

// Count returns number of occurrences of v.
func (l *List[T]) Count(v T) int {
	c := 0
	for _, it := range l.data {
		if reflect.DeepEqual(it, v) {
			c++
		}
	}
	return c
}

// Remove deletes the first occurrence of v. Returns error if not found.
func (l *List[T]) Remove(v T) error {
	idx, err := l.Index(v)
	if err != nil {
		return err
	}
	l.data = append(l.data[:idx], l.data[idx+1:]...)
	return nil
}

// Pop removes and returns element at index i (default last). Supports negative indices.
func (l *List[T]) Pop(idxOpt ...int) (T, error) {
	var zero T
	n := len(l.data)
	if n == 0 {
		return zero, errors.New("pop from empty list")
	}
	idx := n - 1
	if len(idxOpt) > 0 {
		idx = idxOpt[0]
	}
	idx = l.normIndex(idx)
	if idx < 0 || idx >= len(l.data) {
		return zero, errors.New("index out of range")
	}
	val := l.data[idx]
	l.data = append(l.data[:idx], l.data[idx+1:]...)
	return val, nil
}

// Clear removes all elements.
func (l *List[T]) Clear() { l.data = l.data[:0] }

// Reverse reverses the list in-place.
func (l *List[T]) Reverse() {
	for i, j := 0, len(l.data)-1; i < j; i, j = i+1, j-1 {
		l.data[i], l.data[j] = l.data[j], l.data[i]
	}
}

// Sort sorts the list in-place using provided less function. If less is nil returns error.
func (l *List[T]) Sort(less func(a, b T) bool) error {
	if less == nil {
		return errors.New("nil less function")
	}
	sort.Slice(l.data, func(i, j int) bool { return less(l.data[i], l.data[j]) })
	return nil
}

// Copy returns a shallow copy of the list.
func (l *List[T]) Copy() *List[T] {
	return ListFromSlice(l.data)
}

// Slice returns a new List corresponding to l[start:stop]. Supports negative indices and nil-like behavior.
func (l *List[T]) Slice(start, stop int) *List[T] {
	n := len(l.data)
	if start < 0 {
		start = n + start
	}
	if stop < 0 {
		stop = n + stop
	}
	if start < 0 {
		start = 0
	}
	if stop > n {
		stop = n
	}
	if start > stop {
		return NewList[T]()
	}
	d := make([]T, stop-start)
	copy(d, l.data[start:stop])
	return &List[T]{data: d}
}

// Contains returns true if v is present.
func (l *List[T]) Contains(v T) bool {
	_, err := l.Index(v)
	return err == nil
}
