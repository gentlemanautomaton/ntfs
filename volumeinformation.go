package ntfs

import (
	"encoding/binary"
	"fmt"

	"github.com/gentlemanautomaton/ntfs/volumeflag"
)

// https://flatcap.org/linux-ntfs/ntfs/attributes/volume_information.html

// VolumeInformationLength is the length of a volume information record
// in bytes.
const VolumeInformationLength = 12

// VolumeInformation holds volume information about an NTFS volume.
type VolumeInformation struct {
	reserved1    uint64
	VersionMajor uint8
	VersionMinor uint8
	Flags        volumeflag.Flag
}

// UnmarshalBinary unmarshals the little-endian binary representation
// of a volume information record into info.
//
// The provided data must be at least 16 bytes long.
func (info *VolumeInformation) UnmarshalBinary(data []byte) error {
	if len(data) < VolumeInformationLength {
		return ErrTruncatedData
	}
	info.reserved1 = binary.LittleEndian.Uint64(data[0:8])
	info.VersionMajor = data[8]
	info.VersionMinor = data[9]
	info.Flags = volumeflag.Unmarshal(data[10:12])
	return nil
}

// String returns a description of the volume information.
func (info *VolumeInformation) String() string {
	if info.Flags == 0 {
		return fmt.Sprintf("NTFS v%d.%d", info.VersionMajor, info.VersionMinor)
	}
	return fmt.Sprintf("NTFS v%d.%d (Flags: %s)", info.VersionMajor, info.VersionMinor, info.Flags)
}
