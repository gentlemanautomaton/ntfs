package ntfs

import "io"

// Reader is an NTFS file system reader that supports NTFS file system versions
// 3.0 and 3.1. It reads data from an underlying io.ReadSeeker that must not
// include partition table data.
type Reader struct {
	r io.ReadSeeker
}

// NewReader returns a new NTFS filesystem reader that reads from
func NewReader(reader io.ReadSeeker) *Reader {
	return &Reader{
		r: reader,
	}
}

// Reload causes the reader to dismiss its cached data and re-read the file
// system metadata.
func reload() {
}
