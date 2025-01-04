package util

func Ptr[T any](value T) *T {
	return &value
}
