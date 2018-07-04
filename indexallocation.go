package ntfs

// IndexAllocationMinLength is the minimum length of an $INDEX_ALLOCATION
// attribute in bytes.
const IndexAllocationMinLength = 66

// IndexAllocation stores $INDEX_ALLOCATION attribute information.
//
// https://msdn.microsoft.com/library/bb470123
type IndexAllocation struct {
	ParentDirectory FileReference
	reserved        [56]byte // Reserved
	FileNameLength  uint8
	Value           string
}

// UnmarshalBinary unmarshals the little-endian binary representation
// of an $INDEX_ALLOCATION attribute into index.
func (index *IndexAllocation) UnmarshalBinary(data []byte) error {
	return nil
}
