package attrflag

// Flag is an attribute flag.
type Flag uint16

// Attribute flags.
const (
	CompressionMask Flag = 0x00FF // ATTRIBUTE_FLAG_COMPRESSION_MASK
	Encrypted       Flag = 0x4000 // ATTRIBUTE_FLAG_ENCRYPTED
	Sparse          Flag = 0x8000 // ATTRIBUTE_FLAG_SPARSE
)
