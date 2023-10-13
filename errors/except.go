package errors

import "errors"

func Except(err error, ignore ...error) {
	if err == nil {
		return
	}

	for _, e := range ignore {
		if errors.Is(err, e) {
			return
		}
	}

	callback.fn(err)
}

func ExceptFn[W W1](fn W, ignore ...error) {
	Except(fn(), ignore...)
}

func ExceptFn2[T any, W W2[T]](fn W, ignore ...error) T {
	t, err := fn()
	Except(err, ignore...)
	return t
}

func ExceptFn3[T, U any, W W3[T, U]](fn W, ignore ...error) (T, U) {
	t, u, err := fn()
	Except(err, ignore...)
	return t, u
}
