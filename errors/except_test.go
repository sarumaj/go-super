package errors

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestExcept(t *testing.T) {
	type args struct {
		err    error
		ignore []error
	}
	for _, tt := range []struct {
		name string
		args args
		want error
	}{
		{"test#1", args{nil, nil}, fmt.Errorf("")},
		{"test#2", args{os.ErrExist, nil}, os.ErrExist},
		{"test#3", args{os.ErrExist, []error{os.ErrExist}}, fmt.Errorf("")},
	} {
		t.Run(tt.name, func(t *testing.T) {
			buffer := bytes.NewBuffer(nil)
			RegisterCallback(func(err error) { buffer.WriteString(err.Error()) })
			Except(tt.args.err, tt.args.ignore...)
			got := buffer.String()

			if got != tt.want.Error() {
				t.Errorf(`Except(%v, %v) failed: got: %q, want: %q`, tt.args.err, tt.args.ignore, got, tt.want)
			}
		})
	}
}

func TestExceptFn(t *testing.T) {
	type args struct {
		fn     func() error
		ignore []error
	}
	for _, tt := range []struct {
		name string
		args args
		want string
	}{
		{"test#1", args{func() error { return nil }, nil}, ""},
		{"test#2", args{func() error { return os.ErrExist }, nil}, os.ErrExist.Error()},
		{"test#3", args{func() error { return os.ErrExist }, []error{os.ErrExist}}, ""},
	} {
		t.Run(tt.name, func(t *testing.T) {
			buffer := bytes.NewBuffer(nil)
			RegisterCallback(func(err error) { buffer.WriteString(err.Error()) })
			ExceptFn(tt.args.fn, tt.args.ignore...)
			got := buffer.String()

			if got != tt.want {
				t.Errorf(`ExceptFn(%p, %v) failed: got: %q, want: %q`, tt.args.fn, tt.args.ignore, got, tt.want)
			}
		})
	}
}

func TestExceptFn2(t *testing.T) {
	type args struct {
		fn     func() (any, error)
		ignore []error
	}
	for _, tt := range []struct {
		name string
		args args
		want string
	}{
		{"test#1", args{func() (any, error) { return &struct{}{}, nil }, nil}, ""},
		{"test#2", args{func() (any, error) { return nil, os.ErrExist }, nil}, os.ErrExist.Error()},
		{"test#3", args{func() (any, error) { return &struct{}{}, os.ErrExist }, []error{os.ErrExist}}, ""},
	} {
		t.Run(tt.name, func(t *testing.T) {
			buffer := bytes.NewBuffer(nil)
			RegisterCallback(func(err error) { buffer.WriteString(err.Error()) })
			ret := ExceptFn2(tt.args.fn, tt.args.ignore...)
			got := buffer.String()

			if got != tt.want {
				t.Errorf(`ExceptFn(%p, %v) failed: got: %q, want: %q`, tt.args.fn, tt.args.ignore, got, tt.want)
			}

			if got == "" && ret == nil {
				t.Errorf(`ExceptFn(%p, %v) failed: got nil`, tt.args.fn, tt.args.ignore)
			}
		})
	}
}

func TestExceptFn3(t *testing.T) {
	type args struct {
		fn     func() (any, any, error)
		ignore []error
	}
	for _, tt := range []struct {
		name string
		args args
		want string
	}{
		{"test#1", args{func() (any, any, error) { return &struct{}{}, &struct{}{}, nil }, nil}, ""},
		{"test#2", args{func() (any, any, error) { return nil, nil, os.ErrExist }, nil}, os.ErrExist.Error()},
		{"test#3", args{func() (any, any, error) { return &struct{}{}, &struct{}{}, os.ErrExist }, []error{os.ErrExist}}, ""},
	} {
		t.Run(tt.name, func(t *testing.T) {
			buffer := bytes.NewBuffer(nil)
			RegisterCallback(func(err error) { buffer.WriteString(err.Error()) })
			ret1, ret2 := ExceptFn3(tt.args.fn, tt.args.ignore...)
			got := buffer.String()

			if got != tt.want {
				t.Errorf(`ExceptFn(%p, %v) failed: got: %q, want: %q`, tt.args.fn, tt.args.ignore, got, tt.want)
			}

			if got == "" && (ret1 == nil || ret2 == nil) {
				t.Errorf(`ExceptFn(%p, %v) failed: got nil`, tt.args.fn, tt.args.ignore)
			}
		})
	}
}
