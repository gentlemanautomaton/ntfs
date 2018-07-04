package datarun

// Header describes the allocation of bytes in a mapping pair.
type Header byte

// LengthBytes returns the number of length bytes in a mapping pair.
func (h Header) LengthBytes() int {
	return int(h & 0x0f) // Lower nibble
}

// OffsetBytes returns the number of offset bytes in a mapping pair.
func (h Header) OffsetBytes() int {
	return int((h & 0xf0) >> 4) // Upper nibble
}

// DataBytes returns the number of data bytes in a mapping pair,
// excluding the header itself.
func (h Header) DataBytes() int {
	return h.LengthBytes() + h.OffsetBytes()
}

// Valid returns false if the header contains invalid data.
func (h Header) Valid() bool {
	length := h.LengthBytes()
	offset := h.OffsetBytes()
	if length > 8 || offset > 8 {
		// We've exceeded a 64-bit integer
		return false
	}
	if offset > 0 && length == 0 {
		// Length must be present when Offset is present
		return false
	}
	return true
}
