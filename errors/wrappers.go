package errors

type (
	ErrorFn            func() error
	ErrorFn2[T any]    func() (T, error)
	ErrorFn3[T, U any] func() (T, U, error)
)

func W(err error) ErrorFn { return func() error { return err } }

func W2[T any](t T, err error) ErrorFn2[T] { return func() (T, error) { return t, err } }

func W3[T, U any](t T, u U, err error) ErrorFn3[T, U] {
	return func() (T, U, error) { return t, u, err }
}
