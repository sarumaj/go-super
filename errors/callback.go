package errors

import (
	"fmt"
	"os"
)

var callback = (&defaultCallback{}).reset()

type defaultCallback struct {
	fn Callback
}

func (fn *defaultCallback) reset() *defaultCallback {
	fn.fn = func(err error) {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}

	return fn
}

type Callback func(error)

func RestoreCallback() {
	callback.reset()
}

func RegisterCallback(fn Callback) {
	callback.fn = fn
}
