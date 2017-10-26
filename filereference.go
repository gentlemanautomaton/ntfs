package ntfs

// SegmentReference stores an address in the master file table.
//
// https://msdn.microsoft.com/library/bb470211
type SegmentReference struct {
	SegmentNumberLowPart  uint32
	SegmentNumberHighPart uint16
	SequenceNumber        uint16
}

// FileReference is a reference to a file in the master file table.
type FileReference = SegmentReference
