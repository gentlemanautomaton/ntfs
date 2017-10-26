package ntfs

// FileName stores file name attribute information.
//
// https://msdn.microsoft.com/library/bb470123
type FileName struct {
	ParentDirectory FileReference
	_               [0x38]byte // Reserved
	FileNameLength  uint16
	Flags           uint16
	FileName        [1]uint16 // UTF-16 array of varying length
}
