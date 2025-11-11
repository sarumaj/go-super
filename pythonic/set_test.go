package pythonic

import "testing"

func TestSetBasicOperations(t *testing.T) {
	s := NewSet(1, 2, 3)
	if s.Len() != 3 {
		t.Fatalf("expected len 3, got %d", s.Len())
	}

	s.Add(4)
	if !s.Contains(4) {
		t.Fatalf("expected set to contain 4")
	}

	if err := s.Remove(2); err != nil {
		t.Fatalf("remove failed: %v", err)
	}
	if s.Contains(2) {
		t.Fatalf("2 should have been removed")
	}

	s.Discard(999) // should not panic or error

	// Pop until empty
	for s.Len() > 0 {
		_, err := s.Pop()
		if err != nil {
			t.Fatalf("unexpected pop error: %v", err)
		}
	}
	if _, err := s.Pop(); err == nil {
		t.Fatalf("expected pop from empty set to error")
	}
}

func TestSetAlgebra(t *testing.T) {
	a := NewSet(1, 2, 3)
	b := NewSet(3, 4)

	u := a.Union(b)
	if !u.Contains(1) || !u.Contains(4) || u.Len() != 4 {
		t.Fatalf("union incorrect: %v", u.AsSlice())
	}

	i := a.Intersection(b)
	if i.Len() != 1 || !i.Contains(3) {
		t.Fatalf("intersection incorrect: %v", i.AsSlice())
	}

	d := a.Difference(b)
	if d.Contains(3) || d.Len() != 2 {
		t.Fatalf("difference incorrect: %v", d.AsSlice())
	}

	sd := a.SymmetricDifference(b)
	// symmetric diff should be {1,2,4}
	if sd.Len() != 3 || sd.Contains(3) {
		t.Fatalf("symmetric difference incorrect: %v", sd.AsSlice())
	}

	// subset / superset / equal
	c := NewSet(1, 2)
	if !c.IsSubset(a) || !a.IsSuperset(c) {
		t.Fatalf("subset/superset incorrect")
	}

	cc := a.Copy()
	if !cc.Equal(a) {
		t.Fatalf("copy should be equal to original")
	}
	cc.Add(999)
	if cc.Equal(a) {
		t.Fatalf("sets should differ after mutation of copy")
	}
}

func TestSetInPlaceOps(t *testing.T) {
	a := NewSet(1, 2, 3)
	b := NewSet(3, 4)

	// UnionWith
	a.UnionWith(b)
	if a.Len() != 4 || !a.Contains(4) {
		t.Fatalf("UnionWith failed: %v", a.AsSlice())
	}

	// IntersectionWith
	a = NewSet(1, 2, 3)
	a.IntersectionWith(b)
	if a.Len() != 1 || !a.Contains(3) {
		t.Fatalf("IntersectionWith failed: %v", a.AsSlice())
	}

	// DifferenceWith
	a = NewSet(1, 2, 3)
	a.DifferenceWith(b)
	if a.Contains(3) || a.Len() != 2 {
		t.Fatalf("DifferenceWith failed: %v", a.AsSlice())
	}

	// SymmetricDifferenceWith
	a = NewSet(1, 2, 3)
	a.SymmetricDifferenceWith(b)
	// symmetric diff should be {1,2,4}
	if a.Len() != 3 || a.Contains(3) {
		t.Fatalf("SymmetricDifferenceWith failed: %v", a.AsSlice())
	}

	// IsDisjoint
	x := NewSet(10, 11)
	y := NewSet(12)
	if !x.IsDisjoint(y) {
		t.Fatalf("IsDisjoint failed for disjoint sets")
	}
	y.Add(11)
	if x.IsDisjoint(y) {
		t.Fatalf("IsDisjoint failed when sets share element")
	}
}
