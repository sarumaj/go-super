package pythonic

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Accumulate returns running accumulation of elements using fn. For empty input returns nil.
// The first element of result is s[0].
func Accumulate[T any](s []T, fn func(T, T) T) []T {
	if len(s) == 0 {
		return nil
	}
	out := make([]T, 0, len(s))
	acc := s[0]
	out = append(out, acc)
	for _, v := range s[1:] {
		acc = fn(acc, v)
		out = append(out, acc)
	}
	return out
}

// Chain concatenates multiple slices into one.
func Chain[T any](slices ...[]T) []T {
	var total int
	for _, s := range slices {
		total += len(s)
	}
	if total == 0 {
		return nil
	}
	out := make([]T, 0, total)
	for _, s := range slices {
		out = append(out, s...)
	}
	return out
}

// Cycle repeats the slice count times. If count <= 0 returns nil.
func Cycle[T any](s []T, count int) []T {
	if len(s) == 0 || count <= 0 {
		return nil
	}
	out := make([]T, 0, len(s)*count)
	for i := 0; i < count; i++ {
		out = append(out, s...)
	}
	return out
}

// DropWhile drops leading elements satisfying pred and returns the remainder.
func DropWhile[T any](s []T, pred func(T) bool) []T {
	if len(s) == 0 {
		return nil
	}
	i := 0
	for ; i < len(s); i++ {
		if !pred(s[i]) {
			break
		}
	}
	out := make([]T, len(s)-i)
	copy(out, s[i:])
	return out
}

// Islice returns s[start:stop:step] behavior (stop exclusive). Step must not be 0.
func Islice[T any](s []T, start, stop, step int) []T {
	n := len(s)
	if n == 0 || step == 0 {
		return nil
	}
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
	if start >= stop {
		return nil
	}
	out := make([]T, 0, (stop-start+abs(step)-1)/abs(step))
	if step > 0 {
		for i := start; i < stop; i += step {
			out = append(out, s[i])
		}
	} else {
		for i := start; i > stop; i += step {
			out = append(out, s[i])
		}
	}
	return out
}

// Repeat returns a slice containing value repeated count times. If count <=0 returns nil.
func Repeat[T any](value T, count int) []T {
	if count <= 0 {
		return nil
	}
	out := make([]T, count)
	for i := 0; i < count; i++ {
		out[i] = value
	}
	return out
}

// TakeWhile returns the leading slice of elements satisfying pred.
func TakeWhile[T any](s []T, pred func(T) bool) []T {
	if len(s) == 0 {
		return nil
	}
	i := 0
	for ; i < len(s); i++ {
		if !pred(s[i]) {
			break
		}
	}
	if i == 0 {
		return nil
	}
	out := make([]T, i)
	copy(out, s[:i])
	return out
}
