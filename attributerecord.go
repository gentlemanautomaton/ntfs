package ntfs

import "github.com/gentlemanautomaton/ntfs/attrtype"

// AttributeRecordHeader holds information about an attribute record, which may
// be resident or non-resident.
//
// https://msdn.microsoft.com/library/bb470039
type AttributeRecordHeader struct {
	TypeCode     attrtype.Code
	RecordLength uint32
	FormCode     byte
	NameLength   byte
	NameOffset   uint16
	Flags        uint16
	Instance     uint16
	// TODO: Handle Union Data
}
