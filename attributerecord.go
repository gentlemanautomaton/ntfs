package ntfs

import (
	"encoding/binary"

	"github.com/gentlemanautomaton/ntfs/attrform"
	"github.com/gentlemanautomaton/ntfs/attrtype"
)

// https://blogs.technet.microsoft.com/askcore/2009/10/16/the-four-stages-of-ntfs-file-growth/
// https://blogs.technet.microsoft.com/askcore/2010/08/25/ntfs-file-attributes/

// AttributeRecordHeaderLength is the length of an attribute header in bytes,
// excluding its resident or non-resident portion.
const AttributeRecordHeaderLength = 16

// AttributeRecordHeader holds information about an attribute record, which may
// be resident or non-resident.
//
// https://msdn.microsoft.com/library/bb470039
type AttributeRecordHeader struct {
	TypeCode     attrtype.Code
	RecordLength uint32
	FormCode     attrform.Code
	NameLength   uint8
	NameOffset   uint16
	Flags        uint16
	Instance     uint16
}

// Resident returns true if the attribute record header precedes resident
// attribute data.
func (header *AttributeRecordHeader) Resident() bool {
	return header.FormCode == attrform.Resident
}

// UnmarshalBinary unmarshals the little-endian binary representation
// of an attribute record header into header.
//
// The provided data must be at least 16 bytes long.
func (header *AttributeRecordHeader) UnmarshalBinary(data []byte) error {
	/*
		if len(data) < 4 {
			return ErrTruncatedData
		}
		header.TypeCode = attrtype.Code(binary.LittleEndian.Uint32(data[0:4]))
		if header.TypeCode == 0xffffffff {
			return nil
		}
	*/
	if len(data) < AttributeRecordHeaderLength {
		return ErrTruncatedData
	}
	header.TypeCode = attrtype.Unmarshal(data[0:4])
	header.RecordLength = binary.LittleEndian.Uint32(data[4:8])
	header.FormCode = attrform.Code(data[8])
	header.NameLength = data[9]
	header.NameOffset = binary.LittleEndian.Uint16(data[10:12])
	header.Flags = binary.LittleEndian.Uint16(data[12:14])
	header.Instance = binary.LittleEndian.Uint16(data[14:16])
	return nil
}

// ResidentAttributeRecordHeaderLength is the length of the resident
// portion of an attribute header in bytes.
const ResidentAttributeRecordHeaderLength = 8

// ResidentAttributeRecordHeader holds information about the resident
// form of an attribute record.
type ResidentAttributeRecordHeader struct {
	ValueLength uint32
	ValueOffset uint16
	reserved    [2]byte
}

// UnmarshalBinary unmarshals the little-endian binary representation
// of an attribute record header into header.
//
// The provided data must be at least 8 bytes long.
func (header *ResidentAttributeRecordHeader) UnmarshalBinary(data []byte) error {
	if len(data) < ResidentAttributeRecordHeaderLength {
		return ErrTruncatedData
	}
	header.ValueLength = binary.LittleEndian.Uint32(data[0:4])
	header.ValueOffset = binary.LittleEndian.Uint16(data[4:6])
	header.reserved[0] = data[6]
	header.reserved[1] = data[7]
	return nil
}

// NonresidentAttributeRecordHeaderLength is the length of the non-resident
// portion of an attribute header in bytes.
const NonresidentAttributeRecordHeaderLength = 48

// NonresidentAttributeRecordHeader holds information about the non-resident
//  form of an attribute record.
type NonresidentAttributeRecordHeader struct {
	LowestVCN          VCN
	HighestVCN         VCN
	MappingPairsOffset uint16
	reserved           [6]byte
	AllocatedLength    int64
	FileSize           int64
	ValidDataLength    int64
	TotalAllocated     int64
}

// UnmarshalBinary unmarshals the little-endian binary representation
// of the non-resident portion of an attribute record header into header.
//
// The provided data must be at least 8 bytes long.
func (header *NonresidentAttributeRecordHeader) UnmarshalBinary(data []byte) error {
	if len(data) < NonresidentAttributeRecordHeaderLength {
		return ErrTruncatedData
	}
	header.LowestVCN = VCN(binary.LittleEndian.Uint32(data[0:4]))
	header.HighestVCN = VCN(binary.LittleEndian.Uint32(data[4:8]))
	header.MappingPairsOffset = binary.LittleEndian.Uint16(data[8:10])
	header.reserved[0] = data[10]
	header.reserved[1] = data[11]
	header.reserved[2] = data[12]
	header.reserved[3] = data[13]
	header.reserved[4] = data[14]
	header.reserved[5] = data[15]
	header.AllocatedLength = int64(binary.LittleEndian.Uint64(data[16:24]))
	header.FileSize = int64(binary.LittleEndian.Uint64(data[24:32]))
	header.ValidDataLength = int64(binary.LittleEndian.Uint64(data[32:40]))
	header.TotalAllocated = int64(binary.LittleEndian.Uint64(data[40:48]))
	return nil
}
