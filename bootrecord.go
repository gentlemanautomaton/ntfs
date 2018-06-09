package ntfs

import (
	"bytes"
	"io"
)

// BootRecordLength is the length of a volume boot record in bytes.
const BootRecordLength = 84

// BootRecord is the NTFS volume boot record present in the first sector
// of a volume.
type BootRecord struct {
	JumpInstruction [3]byte
	OEMID           [8]byte
	ParameterBlock
}

// ReadFrom reads 84 bytes of volume boot record data from r into boot.
func (boot *BootRecord) ReadFrom(r io.Reader) (n int64, err error) {
	var buf [BootRecordLength]byte
	n32, err := r.Read(buf[:])
	n = int64(n32)
	if err != nil {
		return
	}

	if n != BootRecordLength {
		return 0, ErrTruncatedData
	}

	return n, boot.UnmarshalBinary(buf[:])
}

// UnmarshalBinary unmarshals the little-endian binary representation of
// volume boot record data into boot.
//
// The provided data must be at least 84 bytes long.
func (boot *BootRecord) UnmarshalBinary(data []byte) error {
	if len(data) < BootRecordLength {
		return ErrTruncatedData
	}
	if !bytes.Equal(data[3:11], Label[:]) {
		return ErrInvalidLabel
	}
	copy(boot.JumpInstruction[:], data[0:3])
	copy(boot.OEMID[:], data[3:11])
	return boot.ParameterBlock.UnmarshalBinary(data[11:])
}
