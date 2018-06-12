package filenameflag

import (
	"fmt"
	"strings"
)

// Flag is a file name flag.
type Flag uint8

// File name flags.
const (
	NTFS        Flag = 0x01 // FILE_NAME_NTFS
	DOS         Flag = 0x02 // FILE_NAME_DOS
	KnownMask   Flag = NTFS | DOS
	UnknownMask Flag = ^KnownMask
)

// String returns a description of the file name flags.
func (f Flag) String() string {
	var flags []string

	// Report known flags
	if f&NTFS != 0 {
		flags = append(flags, "NTFS")
	}
	if f&DOS != 0 {
		flags = append(flags, "DOS")
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
