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

	// ErrTruncatedData is returned when attempting to read data from a buffer
	// or reader with insufficient data.
	ErrTruncatedData = errors.New("insufficient or truncated data")

	// ErrInvalidUnicode is returned when attempting to decode invalid unicode
	// data. This happens when the unicode data being processed has an odd
	// number of bytes, exceeds a specified maximum length, or is located
	// beyond the bounds of its containing record.
	ErrInvalidUnicode = errors.New("invalid unicode data")

	// ErrFileNameOutOfBounds is returned when a the value of a file name attribute
	// exceeds the bounds of its containing record.
	//
	// This typically is indicative of MFT corruption.
	ErrFileNameOutOfBounds = errors.New("file name value exceeds the bounds of its file record segment")
)
