package ntfs

// FileRecordSegmentHeader stores file record segment information. It is
// immediately followed by the update sequence array.
//
// https://msdn.microsoft.com/library/bb470124
type FileRecordSegmentHeader struct {
	MultiSectorHeader
	_                     uint64 // reserved
	SequenceNumber        uint16
	_                     uint16 // reserved
	FirstAttributeOffset  uint16
	Flags                 uint16
	_                     [2]uint32 // reserved
	BaseFileRecordSegment FileReference
	_                     [4]uint16 // reserved
}
