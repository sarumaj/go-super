package pythonic

import "testing"

func TestChainAccumulateTakeDropRepeatCycleIslice(t *testing.T) {
	a := []int{1, 2}
	b := []int{3, 4}
	c := Chain(a, b)
	if len(c) != 4 || c[0] != 1 || c[3] != 4 {
		t.Fatalf("Chain failed: %v", c)
	}

	acc := Accumulate([]int{1, 2, 3}, func(x, y int) int { return x + y })
	if len(acc) != 3 || acc[0] != 1 || acc[2] != 6 {
		t.Fatalf("Accumulate failed: %v", acc)
	}

	s := []int{2, 4, 6, 1}
	tw := TakeWhile(s, func(x int) bool { return x%2 == 0 })
	if len(tw) != 3 {
		t.Fatalf("TakeWhile failed: %v", tw)
	}
	dw := DropWhile(s, func(x int) bool { return x%2 == 0 })
	if len(dw) != 1 || dw[0] != 1 {
		t.Fatalf("DropWhile failed: %v", dw)
	}

	rep := Repeat(5, 3)
	if len(rep) != 3 || rep[0] != 5 || rep[2] != 5 {
		t.Fatalf("Repeat failed: %v", rep)
	}

	cyc := Cycle([]int{7, 8}, 2)
	if len(cyc) != 4 || cyc[0] != 7 || cyc[3] != 8 {
		t.Fatalf("Cycle failed: %v", cyc)
	}

	// Islice: positive step
	arr := []int{0, 1, 2, 3, 4, 5}
	sl := Islice(arr, 1, 6, 2)
	if len(sl) != 3 || sl[0] != 1 || sl[2] != 5 {
		t.Fatalf("Islice failed: %v", sl)
	}
}
