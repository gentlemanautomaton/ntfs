package datarun

import "errors"

var (
	// ErrTruncatedData is returned when attempting to read data from a buffer
	// or reader with insufficient data.
	ErrTruncatedData = errors.New("insufficient or truncated data")

	// ErrInvalidHeader is returned when attempting to decode a data run
	// header that has an invalid size.
	ErrInvalidHeader = errors.New("invalid data run header")
)
