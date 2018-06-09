package ntfs

import "io"

// Reader is an NTFS file system reader that supports NTFS file system versions
// 3.0 and 3.1. It reads data from an underlying io.ReadSeeker that must not
// include partition table data.
type Reader struct {
	r    io.ReadSeeker
	boot BootRecord
}

// NewReader returns a new NTFS filesystem reader that reads from rs.
// It will read the volume boot record before returning. If it cannot
// read the volume boot record from rs, an error will be returned.
func NewReader(rs io.ReadSeeker) (*Reader, error) {
	r := &Reader{
		r: rs,
	}
	_, err := r.boot.ReadFrom(rs)
	return r, err
}

// BootRecord returns a copy of the volume boot record.
func (r *Reader) BootRecord() BootRecord {
	return r.boot
}

// Reload causes the reader to dismiss its cached data and re-read the file
// system metadata.
func reload() {
}
