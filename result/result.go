package result

import (
	"errors"
	"reflect"

	supererrors "github.com/sarumaj/go-super/errors"
)

type Result[O any] struct {
	state State
	fault error
	value O
}

func (r Result[O]) Equals(other Result[O]) bool {
	switch {
	case
		r.state != other.state,
		!(r.fault == other.fault || errors.Is(r.fault, other.fault)),
		!reflect.DeepEqual(r.value, other.value):

		return false
	}

	return true
}

func (r Result[O]) Error() error {
	return r.fault
}

func (r Result[O]) IsSuccess() bool {
	return r.state&Success == Success
}

func (r Result[O]) IsFailure() bool {
	return r.state&Failure == Failure
}

func (r Result[O]) State() State {
	return r.state
}

func (r Result[O]) Output() O {
	return r.value
}

func (r *Result[O]) SetError(fault error) *Result[O] {
	r.fault = fault
	return r
}

func (r *Result[O]) SetState(state State) *Result[O] {
	r.state = state
	return r
}

func (r *Result[O]) SetOutput(value O) *Result[O] {
	r.value = value
	return r
}

func GetResult[T any](fn supererrors.ErrorFn[T], ignore ...error) *Result[T] {
	o, err := fn()
	r := (&Result[T]{}).SetOutput(o).SetError(err)

	if err == nil {
		return r.SetState(Success)
	}

	for _, e := range ignore {
		if errors.Is(err, e) {
			return r.SetState(ExpectedFailure).SetError(nil)
		}
	}

	return r.SetState(Failure)
}
