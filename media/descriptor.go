package media

// Descriptor is a one byte media descriptor as used in file allocation
// tables.
type Descriptor uint8

// Media Descriptors
const (
	HardDisk Descriptor = 0xf8
	RAMDisk  Descriptor = 0xfc
	Floppy   Descriptor = 0xf0 // 1.44 MB floppy disk
)
