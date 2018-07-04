package attrflag

import (
	"encoding/binary"
	"fmt"
	"strings"
)

// Flag is an attribute flag.
type Flag uint16

// Attribute flags.
const (
	CompressionMask Flag = 0x00FF // ATTRIBUTE_FLAG_COMPRESSION_MASK
	Encrypted       Flag = 0x4000 // ATTRIBUTE_FLAG_ENCRYPTED
	Sparse          Flag = 0x8000 // ATTRIBUTE_FLAG_SPARSE
	KnownMask       Flag = CompressionMask | Encrypted | Sparse
	UnknownMask     Flag = ^KnownMask
)

// String returns a description of the file name flags.
func (f Flag) String() string {
	var flags []string

	// Report known flags
	if f&CompressionMask != 0 {
		flags = append(flags, "Compressed")
	}
	if f&Encrypted != 0 {
		flags = append(flags, "Encrypted")
	}
	if f&Sparse != 0 {
		flags = append(flags, "Sparse")
	}
	// Report unknown flags
	if f&UnknownMask != 0 {
		for i := uint(0); i < 32; i++ {
			q := Flag(1) << i
			// Find flags that are present
			if q&f == 0 {
				continue
			}
			// Skip flags that we've already identified
			if q&UnknownMask == 0 {
				continue
			}
			flags = append(flags, fmt.Sprintf("%#04x", uint16(q)))
		}
	}

	return strings.Join(flags, ",")
}

// ShortString returns a description of the attribute flags using
// single letter codes.
func (f Flag) ShortString() string {
	c, e, s := " ", " ", " "
	if f&CompressionMask != 0 {
		c = "C"
	}
	if f&Encrypted != 0 {
		e = "E"
	}
	if f&Sparse != 0 {
		s = "S"
	}
	return c + e + s
}

// Unmarshal unmarshals the little-endian binary representation
// of an attribute flag.
//
// The provided data must be at least 2 bytes long, or unmarshal will
// panic.
func Unmarshal(data []byte) Flag {
	return Flag(binary.LittleEndian.Uint16(data[0:2]))
}
