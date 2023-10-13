package errors

type W1 interface {
	func() error
}

type W2[T any] interface {
	func() (T, error)
}

type W3[T, U any] interface {
	func() (T, U, error)
}
