package ntfs

import (
	"github.com/gentlemanautomaton/ntfs/filenameflag"
)

// FileNameHeaderLength is the length of a file name header in bytes.
const FileNameHeaderLength = 66

// FileName stores file name attribute information.
//
// https://msdn.microsoft.com/library/bb470123
type FileName struct {
	ParentDirectory FileReference
	reserved        [56]byte // Reserved
	FileNameLength  uint8
	Flags           filenameflag.Flag
	Value           string
}

// UnmarshalBinary unmarshals the little-endian binary representation
// of a file name record into entry.
func (entry *FileName) UnmarshalBinary(data []byte) error {
	if len(data) < FileNameHeaderLength {
		return ErrTruncatedData
	}
	if err := entry.ParentDirectory.UnmarshalBinary(data[0:8]); err != nil {
		return err
	}
	copy(entry.reserved[:], data[8:64])
	entry.FileNameLength = uint8(data[64])
	entry.Flags = filenameflag.Flag(data[65])
	start := 66
	length := int(entry.FileNameLength) * 2 // Length is in characters, assuming 16-bit unicode
	end := start + length
	if end > len(data) {
		return ErrFileNameOutOfBounds
	}
	var err error
	entry.Value, err = utf16ToString(data[start:end])
	return err
}
