package ntfs

import "github.com/gentlemanautomaton/ntfs/attrtype"

// AttributeListEntry stores an entry in an attribute list.
type AttributeListEntry struct {
	TypeCode            attrtype.Code
	RecordLength        uint16
	AttributeNameLength uint16
	AttributeNameOffset uint16
	LowestVCN           VCN
	SegmentReference    SegmentReference
	_                   uint16    // Reserved
	AttributeName       [1]uint16 // UTF-16 array of varying length
}
