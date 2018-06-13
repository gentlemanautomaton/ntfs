package ntfs

import (
	"bytes"
	"strconv"
)

// Bitmap stores a series of bits.
type Bitmap []byte

// UnmarshalBinary unmarshals the little-endian binary representation
// of a bitmap attributre into bitmap.
func (bitmap *Bitmap) UnmarshalBinary(data []byte) error {
	*bitmap = Bitmap(data)
	return nil
}

// String returns a description of the bitmap.
func (bitmap Bitmap) String() string {
	if len(bitmap) == 0 {
		return ""
	}
	const limit = 32
	var buf bytes.Buffer
	buf.WriteString("0b")
	for i := range bitmap {
		if i > limit {
			buf.WriteString("...")
			break
		}
		buf.WriteString(strconv.FormatUint(uint64(bitmap[i]), 2))
	}
	return buf.String()
}
