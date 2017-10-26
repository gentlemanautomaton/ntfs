package ntfs

// MultiSectorHeader specifies the location and size of an update sequence
// array.
//
// https://msdn.microsoft.com/library/bb470212
type MultiSectorHeader struct {
	Signature                 [4]byte // Always "FILE"
	UpdateSequenceArrayOffset uint16
	UpdateSequenceArraySize   uint16
}
