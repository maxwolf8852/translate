package translate

import "errors"

var (
	ErrStatusCode      = errors.New("incorrect status code")
	ErrNotFound        = errors.New("not found")
	ErrUnknownProvider = errors.New("unknown provider")
)
