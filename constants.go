package ntfs

import "errors"

var (
	// Label is the OEM ID of NTFS file system boot records.
	Label = [8]byte{'N', 'T', 'F', 'S', ' ', ' ', ' ', ' '}
)

var (
	// ErrInvalidLabel is returned when creating a new reader with invalid
	// NTFS data.
	ErrInvalidLabel = errors.New("volume boot record does not contain a valid NTFS file system label")
)

var (
	// ErrTruncatedData is returned when attempting to read data from a buffer
	// or reader with insufficient data.
	ErrTruncatedData = errors.New("insufficient or truncated data")
)
