package volumeflag

import (
	"encoding/binary"
	"fmt"
	"strings"
)

// https://flatcap.org/linux-ntfs/ntfs/attributes/volume_information.html

// Flag is a volume information flag.
type Flag uint16

// Volume information flags.
const (
	Dirty             Flag = 0x0001
	ResizeLogFile     Flag = 0x0002
	UpgradeOnMount    Flag = 0x0004
	MountedOnNT4      Flag = 0x0008
	DeleteUSNUnderway Flag = 0x0010
	RepairObjectID    Flag = 0x0020
	ChkdskUnderway    Flag = 0x4000
	ModifiedByChkdsk  Flag = 0x8000
	KnownMask         Flag = Dirty | ResizeLogFile | UpgradeOnMount | MountedOnNT4 | DeleteUSNUnderway | RepairObjectID | ChkdskUnderway | ModifiedByChkdsk
	UnknownMask       Flag = ^KnownMask
)

// Unmarshal unmarshals the little-endian binary representation
// of volume information flags.
//
// The provided data must be at least 2 bytes long, or unmarshal will
// panic.
func Unmarshal(data []byte) Flag {
	return Flag(binary.LittleEndian.Uint16(data[0:2]))
}

// String returns a description of the volume information flags.
func (f Flag) String() string {
	var flags []string

	// Report known flags
	if f&Dirty != 0 {
		flags = append(flags, "Dirty")
	}
	if f&ResizeLogFile != 0 {
		flags = append(flags, "ResizeLogFile")
	}
	if f&UpgradeOnMount != 0 {
		flags = append(flags, "UpgradeOnMount")
	}
	if f&MountedOnNT4 != 0 {
		flags = append(flags, "MountedOnNT4")
	}
	if f&DeleteUSNUnderway != 0 {
		flags = append(flags, "DeleteUSNUnderway")
	}
	if f&RepairObjectID != 0 {
		flags = append(flags, "RepairObjectID")
	}
	if f&ChkdskUnderway != 0 {
		flags = append(flags, "ChkdskUnderway")
	}
	if f&ModifiedByChkdsk != 0 {
		flags = append(flags, "ModifiedByChkdsk")
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
