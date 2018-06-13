package ntfs

import (
	"encoding/binary"
	"fmt"
	"time"
)

// StandardInformationMinLength is the minimum length of a standard
// information attribute in bytes.
const StandardInformationMinLength = 48

// StandardInformation holds standard attribute data.
//
// https://msdn.microsoft.com/library/bb545266
type StandardInformation struct {
	FileCreation       time.Time
	FileModification   time.Time
	MFTModification    time.Time
	FileRead           time.Time
	DOSFilePermissions uint32
	MaxVersions        uint32
	VersionNumber      uint32
	ClassID            uint32
	OwnerID            uint32
	SecurityID         uint32
	QuotaCharged       uint64
	USN                uint64
}

// UnmarshalBinary unmarshals the little-endian binary representation
// of a standard information attribute into info.
//
// The provided data must be at least 48 bytes long.
func (info *StandardInformation) UnmarshalBinary(data []byte) error {
	if len(data) < StandardInformationMinLength {
		return ErrTruncatedData
	}
	info.FileCreation = unmarshalFileTime(data[0:8])
	info.FileModification = unmarshalFileTime(data[8:16])
	info.MFTModification = unmarshalFileTime(data[16:24])
	info.FileRead = unmarshalFileTime(data[24:32])
	info.DOSFilePermissions = binary.LittleEndian.Uint32(data[32:36])
	info.MaxVersions = binary.LittleEndian.Uint32(data[36:40])
	info.VersionNumber = binary.LittleEndian.Uint32(data[40:44])
	info.ClassID = binary.LittleEndian.Uint32(data[44:48])
	if len(data) >= 72 {
		info.OwnerID = binary.LittleEndian.Uint32(data[48:52])
		info.SecurityID = binary.LittleEndian.Uint32(data[52:56])
		info.QuotaCharged = binary.LittleEndian.Uint64(data[56:64])
		info.USN = binary.LittleEndian.Uint64(data[64:72])
	}
	return nil
}

// String returns a description of the standard information.
func (info *StandardInformation) String() string {
	c := localTimeString(info.FileCreation)
	m := localTimeString(info.FileModification)
	r := localTimeString(info.FileRead)
	mft := localTimeString(info.MFTModification)
	output := fmt.Sprintf("Created: %v, Updated: %v, Read: %v, MFT: %v", c, m, r, mft)
	if info.OwnerID != 0 || info.SecurityID != 0 || info.QuotaCharged != 0 || info.USN != 0 {
		output += fmt.Sprintf(", OID: %d, SID: %d, QC: %d, USN: %d",
			info.OwnerID, info.SecurityID, info.QuotaCharged, info.USN)
	}
	return output
}
