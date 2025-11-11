package pythonic

import "testing"

func TestDictBasicOperations(t *testing.T) {
	d := NewDictFromMap(map[string]int{"a": 1, "b": 2})
	if d.Len() != 2 {
		t.Fatalf("expected len 2, got %d", d.Len())
	}

	if !d.Contains("a") {
		t.Fatalf("expected contains a")
	}

	v, err := d.Get("a")
	if err != nil || v != 1 {
		t.Fatalf("expected get a=1, got %v err %v", v, err)
	}

	d.Set("c", 3)
	if !d.Contains("c") {
		t.Fatalf("expected contains c after set")
	}

	if err := d.Remove("b"); err != nil {
		t.Fatalf("remove failed: %v", err)
	}
	if d.Contains("b") {
		t.Fatalf("b should have been removed")
	}

	// Pop existing
	pv, err := d.Pop("a")
	if err != nil || pv != 1 {
		t.Fatalf("pop a failed: %v %v", pv, err)
	}

	// PopItem until empty
	d.Set("x", 10)
	d.Set("y", 20)
	for d.Len() > 0 {
		_, _, err := d.PopItem()
		if err != nil {
			t.Fatalf("popitem failed: %v", err)
		}
	}
}

func TestDictCopyUpdateEqual(t *testing.T) {
	a := NewDictFromMap(map[string]int{"k1": 1})
	b := a.Copy()
	if !a.Equal(b) {
		t.Fatalf("copy must be equal")
	}
	b.Set("k2", 2)
	if a.Equal(b) {
		t.Fatalf("after modifying copy, they should differ")
	}

	a.Update(b)
	if !a.Contains("k2") {
		t.Fatalf("update failed")
	}

	// Keys/Values/Items length checks
	if len(a.Keys()) == 0 || len(a.Values()) == 0 || len(a.Items()) == 0 {
		t.Fatalf("expected keys/values/items non-empty")
	}

	// SetDefault behavior
	dv := a.SetDefault("k3", 99)
	if dv != 99 || !a.Contains("k3") {
		t.Fatalf("SetDefault failed")
	}
}
