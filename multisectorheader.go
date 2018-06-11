package ntfs

import "encoding/binary"

// MultiSectorHeaderLength is the length of a multi-sector header in bytes.
const MultiSectorHeaderLength = 8

// MultiSectorHeader specifies the location and size of an update sequence
// array.
//
// https://msdn.microsoft.com/library/bb470212
type MultiSectorHeader struct {
	Signature                 [4]byte // Always "FILE"
	UpdateSequenceArrayOffset uint16
	UpdateSequenceArraySize   uint16
}

// UnmarshalBinary unmarshals the little-endian binary representation
// of a mutli-sector header into hdr.
//
// The provided data must be at least 73 bytes long.
func (header *MultiSectorHeader) UnmarshalBinary(data []byte) error {
	if len(data) < MultiSectorHeaderLength {
		return ErrTruncatedData
	}
	header.Signature[0] = data[0]
	header.Signature[1] = data[1]
	header.Signature[2] = data[2]
	header.Signature[3] = data[3]
	header.UpdateSequenceArrayOffset = binary.LittleEndian.Uint16(data[4:6])
	header.UpdateSequenceArraySize = binary.LittleEndian.Uint16(data[6:8])
	return nil
}
