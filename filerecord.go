package ntfs

import (
	"encoding/binary"
	"io"
)

// FileRecordSegmentHeaderLength is the length of a file record segment
// header in bytes.
const FileRecordSegmentHeaderLength = 42

// FileRecordSegmentHeader stores file record segment information. It is
// present at the start of each file record, and is immediately followed
// by the update sequence array.
//
// https://msdn.microsoft.com/library/bb470124
type FileRecordSegmentHeader struct {
	MultiSectorHeader
	logFileSequenceNumber uint64 // reserved
	SequenceNumber        uint16
	hardLinkCount         uint16 // reserved
	FirstAttributeOffset  uint16
	Flags                 uint16
	actualSize            uint32 // reserved
	allocatedSize         uint32 // reserved
	BaseFileRecordSegment FileReference
	reserved4             uint16 // reserved
}

// ReadFrom reads 42 bytes of file record segment header data
// from r into header.
func (header *FileRecordSegmentHeader) ReadFrom(r io.Reader) (n int64, err error) {
	var buf [FileRecordSegmentHeaderLength]byte
	n32, err := r.Read(buf[:])
	n = int64(n32)
	if err != nil {
		return
	}
	if n != FileRecordSegmentHeaderLength {
		return 0, ErrTruncatedData
	}
	return n, header.UnmarshalBinary(buf[:])
}

// UnmarshalBinary unmarshals the little-endian binary representation
// of a file record segment header into header.
//
// The provided data must be at least 42 bytes long.
func (header *FileRecordSegmentHeader) UnmarshalBinary(data []byte) error {
	if len(data) < FileRecordSegmentHeaderLength {
		return ErrTruncatedData
	}
	if err := header.MultiSectorHeader.UnmarshalBinary(data[0:8]); err != nil {
		return err
	}
	header.logFileSequenceNumber = binary.LittleEndian.Uint64(data[8:16])
	header.SequenceNumber = binary.LittleEndian.Uint16(data[16:18])
	header.hardLinkCount = binary.LittleEndian.Uint16(data[18:20])
	header.FirstAttributeOffset = binary.LittleEndian.Uint16(data[20:22])
	header.Flags = binary.LittleEndian.Uint16(data[22:24])
	header.actualSize = binary.LittleEndian.Uint32(data[24:28])
	header.allocatedSize = binary.LittleEndian.Uint32(data[28:32])
	if err := header.BaseFileRecordSegment.UnmarshalBinary(data[32:40]); err != nil {
		return err
	}
	header.reserved4 = binary.LittleEndian.Uint16(data[40:42])
	return nil
}
