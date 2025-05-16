package translate

import "errors"

var (
	ErrStatusCode      = errors.New("incorrect status code")
	ErrNotFound        = errors.New("not found")
	ErrNoProviders     = errors.New("no providers specified")
	ErrTooManyRequests = errors.New("too many requests")
	ErrInvalidInput    = errors.New("invalid input")
	ErrUnsupportedLang = errors.New("unsupported language")
)
