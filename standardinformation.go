package ntfs

// StandardInformation holds standard attribute data.
//
// https://msdn.microsoft.com/library/bb545266
type StandardInformation struct {
	_          [0x30]byte // Reserved
	OwnerID    uint32
	SecurityID uint32
}
