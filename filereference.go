package ntfs

import "encoding/binary"

// SegmentReferenceLength is the length of a segment reference in bytes.
const SegmentReferenceLength = 8

// SegmentReference stores an address in the master file table.
//
// https://msdn.microsoft.com/library/bb470211
type SegmentReference struct {
	SegmentNumberLowPart  uint32
	SegmentNumberHighPart uint16
	SequenceNumber        uint16
}

// IsZero returns true if the segment reference is zero.
func (ref *SegmentReference) IsZero() bool {
	return ref.SegmentNumberLowPart == 0 && ref.SegmentNumberHighPart == 0 && ref.SequenceNumber == 0
}

// UnmarshalBinary unmarshals the little-endian binary representation
// of a segment reference into ref.
//
// The provided data must be at least 8 bytes long.
func (ref *SegmentReference) UnmarshalBinary(data []byte) error {
	if len(data) < SegmentReferenceLength {
		return ErrTruncatedData
	}
	ref.SegmentNumberLowPart = binary.LittleEndian.Uint32(data[0:4])
	ref.SegmentNumberHighPart = binary.LittleEndian.Uint16(data[4:6])
	ref.SequenceNumber = binary.LittleEndian.Uint16(data[6:8])
	return nil
}

// FileReference is a reference to a file in the master file table.
type FileReference = SegmentReference
