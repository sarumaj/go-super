package pythonic

import (
	"testing"
)

func TestMapFilterReduceAnyAll(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	mapped := Map(nums, func(x int) int { return x * 2 })
	if len(mapped) != 4 || mapped[0] != 2 || mapped[3] != 8 {
		t.Fatalf("Map produced wrong result: %v", mapped)
	}

	filtered := Filter(nums, func(x int) bool { return x%2 == 0 })
	if len(filtered) != 2 || filtered[0] != 2 || filtered[1] != 4 {
		t.Fatalf("Filter produced wrong result: %v", filtered)
	}

	sum := Reduce(nums, 0, func(acc, v int) int { return acc + v })
	if sum != 10 {
		t.Fatalf("Reduce sum expected 10, got %d", sum)
	}

	if !Any(nums, func(x int) bool { return x == 3 }) {
		t.Fatalf("Any failed")
	}
	if !All(nums, func(x int) bool { return x > 0 }) {
		t.Fatalf("All failed")
	}
}

func TestZipEnumerateRangeInts(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	z := Zip(a, b)
	if len(z) != 3 || len(z[0]) != 2 || z[1][0] != 2 || z[1][1] != 5 {
		t.Fatalf("Zip produced wrong result: %v", z)
	}

	en := Enumerate(a, 10)
	if len(en) != 3 || en[0].Idx != 10 || en[2].Val != 3 {
		t.Fatalf("Enumerate produced wrong result: %v", en)
	}

	r := RangeInts(0, 5, 2)
	if len(r) != 3 || r[0] != 0 || r[2] != 4 {
		t.Fatalf("RangeInts produced wrong result: %v", r)
	}
}

func TestSumMinMaxSorted(t *testing.T) {
	ints := []int{5, 1, 3}
	if Sum(ints) != 9 {
		t.Fatalf("Sum incorrect for ints")
	}
	if v, err := Min(ints); err != nil || v != 1 {
		t.Fatalf("Min incorrect: %v %v", v, err)
	}
	if v, err := Max(ints); err != nil || v != 5 {
		t.Fatalf("Max incorrect: %v %v", v, err)
	}
	sorted := Sorted(ints)
	if len(sorted) != 3 || sorted[0] != 1 || sorted[2] != 5 {
		t.Fatalf("Sorted incorrect: %v", sorted)
	}
}
