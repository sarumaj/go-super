package pythonic

import "testing"

func TestListBasicOperations(t *testing.T) {
	l := NewList(1, 2, 3)
	if l.Len() != 3 {
		t.Fatalf("expected len 3, got %d", l.Len())
	}

	l.Append(4)
	if l.Len() != 4 {
		t.Fatalf("expected len 4 after append, got %d", l.Len())
	}

	v, err := l.Get(-1)
	if err != nil || v != 4 {
		t.Fatalf("expected last element 4, got %v err %v", v, err)
	}

	if err := l.Set(0, 10); err != nil {
		t.Fatalf("set failed: %v", err)
	}
	v0, _ := l.Get(0)
	if v0 != 10 {
		t.Fatalf("expected first element 10, got %v", v0)
	}

	if idx, err := l.Index(2); err != nil || idx != 1 {
		t.Fatalf("expected index of 2 to be 1, got %d err %v", idx, err)
	}

	if c := l.Count(2); c != 1 {
		t.Fatalf("expected count 1, got %d", c)
	}

	if !l.Contains(3) {
		t.Fatalf("expected contains 3")
	}

	if err := l.Remove(2); err != nil {
		t.Fatalf("remove failed: %v", err)
	}
	if l.Contains(2) {
		t.Fatalf("2 should have been removed")
	}

	popped, err := l.Pop()
	if err != nil || popped != 4 {
		t.Fatalf("expected pop 4, got %v err %v", popped, err)
	}

	// pop by index
	p2, err := l.Pop(0)
	if err != nil || p2 != 10 {
		t.Fatalf("expected pop index 0 -> 10, got %v err %v", p2, err)
	}

	l.Clear()
	if l.Len() != 0 {
		t.Fatalf("expected cleared list to have len 0")
	}
}

func TestListInsertExtendSliceReverseSort(t *testing.T) {
	l := NewList(3, 1)
	l.Insert(1, 2)
	if v, _ := l.Get(1); v != 2 {
		t.Fatalf("expected inserted 2 at index 1, got %v", v)
	}

	l.Extend(4, 5)
	if l.Len() != 5 {
		t.Fatalf("expected len 5 after extend, got %d", l.Len())
	}

	s := l.Slice(1, 4)
	if s.Len() != 3 {
		t.Fatalf("expected slice len 3, got %d", s.Len())
	}

	l.Reverse()
	// after reverse, first element should be 5
	if v, _ := l.Get(0); v != 5 {
		t.Fatalf("expected first element 5 after reverse, got %v", v)
	}

	// sort ascending using comparator
	err := l.Sort(func(a, b int) bool { return a < b })
	if err != nil {
		t.Fatalf("sort failed: %v", err)
	}
	// verify sorted
	prev, _ := l.Get(0)
	for i := 1; i < l.Len(); i++ {
		cur, _ := l.Get(i)
		// prev and cur are of type any under generics in tests; assert as int
		if prev > cur {
			t.Fatalf("list not sorted: %v > %v", prev, cur)
		}
		prev = cur
	}

	// copy
	c := l.Copy()
	if c.Len() != l.Len() {
		t.Fatalf("copy length mismatch")
	}
	c.Append(999)
	if c.Len() == l.Len() {
		t.Fatalf("copy should be independent from original")
	}
}

func TestSortNilLessProducesError(t *testing.T) {
	l := NewList(2, 1)
	if err := l.Sort(nil); err == nil {
		t.Fatalf("expected error when sorting with nil less")
	}
}

func TestConcatVariants(t *testing.T) {
	a := NewList(1, 2)
	b := NewList(3, 4)

	// concat with another *List[T]
	c := a.Concat(b)
	if c.Len() != 4 {
		t.Fatalf("expected concat len 4, got %d", c.Len())
	}
	v, _ := c.Get(2)
	if v != 3 {
		t.Fatalf("expected element 3 at index 2, got %v", v)
	}
	// originals unchanged
	if a.Len() != 2 || b.Len() != 2 {
		t.Fatalf("original lists must be unchanged")
	}

	// concat with []T
	s := []int{5, 6}
	c2 := a.Concat(s)
	if c2.Len() != 4 {
		t.Fatalf("expected concat with slice len 4, got %d", c2.Len())
	}
	vv, _ := c2.Get(3)
	if vv != 6 {
		t.Fatalf("expected element 6 at index 3, got %v", vv)
	}

	// concat with nil should produce a copy
	c3 := a.Concat(nil)
	if c3.Len() != a.Len() {
		t.Fatalf("expected concat(nil) to produce same-length copy")
	}
	c3.Append(99)
	if a.Len() == c3.Len() {
		t.Fatalf("modifying concat result must not change original")
	}

	// concat with unsupported type returns a copy (current behavior)
	c4 := a.Concat("unsupported")
	if c4.Len() != a.Len() {
		t.Fatalf("expected concat unsupported to return a copy with same length")
	}
	c4.Append(100)
	if a.Len() == c4.Len() {
		t.Fatalf("modifying concat result from unsupported type must not change original")
	}
}
