package errors

import "errors"

func Except(err error, ignore ...error) {
	if err == nil {
		return
	}

	lastError.store(err)

	for _, e := range ignore {
		if errors.Is(err, e) {
			return
		}
	}

	callback.fn(err)
}

func ExceptFn(fn ErrorFn, ignore ...error) { Except(fn(), ignore...) }

func ExceptFn2[T any](fn ErrorFn2[T], ignore ...error) T {
	t, err := fn()
	Except(err, ignore...)
	return t
}

func ExceptFn3[T, U any](fn ErrorFn3[T, U], ignore ...error) (T, U) {
	t, u, err := fn()
	Except(err, ignore...)
	return t, u
}
