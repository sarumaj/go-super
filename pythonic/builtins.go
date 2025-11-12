package pythonic

import (
	"errors"
	"slices"
)

// EnumeratePair is used by Enumerate to return index and value.
type EnumeratePair[T any] struct {
	Idx int
	Val T
}

// Numeric matches integer and floating-point numeric types.
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Ordered is a constraint that matches ordered types for comparisons and sorting.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

// Any returns true if any element satisfies pred.
func Any[T any](s []T, pred func(T) bool) bool {
	return slices.ContainsFunc(s, pred)
}

// All returns true if all elements satisfy pred (true for empty slice).
func All[T any](s []T, pred func(T) bool) bool {
	for _, v := range s {
		if !pred(v) {
			return false
		}
	}
	return true
}

// Enumerate returns pairs of (index+start, value).
func Enumerate[T any](s []T, start int) []EnumeratePair[T] {
	if s == nil {
		return nil
	}
	out := make([]EnumeratePair[T], 0, len(s))
	for i, v := range s {
		out = append(out, EnumeratePair[T]{Idx: start + i, Val: v})
	}
	return out
}

// Filter returns elements where pred(elem) is true.
func Filter[T any](s []T, pred func(T) bool) []T {
	if s == nil {
		return nil
	}
	out := make([]T, 0, len(s))
	for _, v := range s {
		if pred(v) {
			out = append(out, v)
		}
	}
	return out
}

// Map applies fn to each element and returns a new slice.
func Map[T any, U any](s []T, fn func(T) U) []U {
	if s == nil {
		return nil
	}
	out := make([]U, 0, len(s))
	for _, v := range s {
		out = append(out, fn(v))
	}
	return out
}

// Max returns the maximum element for Ordered types.
func Max[T Ordered](s []T) (T, error) {
	var zero T
	if len(s) == 0 {
		return zero, errors.New("max of empty slice")
	}
	m := s[0]
	for _, v := range s[1:] {
		if v > m {
			m = v
		}
	}
	return m, nil
}

// Min returns the minimum element for Ordered types.
func Min[T Ordered](s []T) (T, error) {
	var zero T
	if len(s) == 0 {
		return zero, errors.New("min of empty slice")
	}
	m := s[0]
	for _, v := range s[1:] {
		if v < m {
			m = v
		}
	}
	return m, nil
}

// RangeInts returns a slice of ints from start (inclusive) to stop (exclusive) with step.
func RangeInts(start, stop, step int) []int {
	if step == 0 {
		return nil
	}
	var out []int
	if step > 0 {
		for i := start; i < stop; i += step {
			out = append(out, i)
		}
	} else {
		for i := start; i > stop; i += step {
			out = append(out, i)
		}
	}
	return out
}

// Reduce reduces slice to a single value using accumulator function.
func Reduce[T any, R any](s []T, init R, fn func(R, T) R) R {
	acc := init
	for _, v := range s {
		acc = fn(acc, v)
	}
	return acc
}

// Sorted returns a sorted copy of the slice for Ordered types.
func Sorted[T Ordered](s []T) []T {
	if s == nil {
		return nil
	}
	out := make([]T, len(s))
	copy(out, s)
	slices.Sort(out)
	return out
}

// Sum sums a slice of any numeric type and returns the same numeric type.
func Sum[T Numeric](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

// SumInts sums a slice of ints.
func SumInts(s []int) int {
	return Sum(s)
}

// SumFloat64 sums a slice of float64s.
func SumFloat64(s []float64) float64 {
	return Sum(s)
}

// Zip zips slices of the same element type into a slice of rows.
// The result length is the length of the shortest input slice.
func Zip[T any](slices ...[]T) [][]T {
	if len(slices) == 0 {
		return nil
	}
	min := len(slices[0])
	for _, s := range slices[1:] {
		if len(s) < min {
			min = len(s)
		}
	}
	res := make([][]T, 0, min)
	for i := 0; i < min; i++ {
		row := make([]T, len(slices))
		for j, s := range slices {
			row[j] = s[i]
		}
		res = append(res, row)
	}
	return res
}
