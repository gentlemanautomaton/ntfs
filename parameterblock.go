package ntfs

import (
	"encoding/binary"
	"io"

	"github.com/gentlemanautomaton/ntfs/media"
)

// https://en.wikipedia.org/wiki/BIOS_parameter_block#NTFS
// https://docs.microsoft.com/en-us/previous-versions/windows/it-pro/windows-server-2003/cc781134(v=ws.10)

// ParameterBlockLength is the length of an NTFS parameter block in bytes.
const ParameterBlockLength = 73

// ParameterBlock is the 73-byte BIOS parameter block that forms part of an
// NFTS boot record.
type ParameterBlock struct {
	// DOS 2.0 parameter block
	BytesPerSector       uint16           //  0:2  The logical block size
	SectorsPerCluster    uint8            //  2:3  The cluster size (must be power of 2 greater than 0)
	reservedSectors      uint16           //  3:5  reserved (must be zero)
	numberOfFATs         uint8            //  5:6  reserved (must be zero)
	rootDirectoryEntries uint16           //  6:8  reserved (must be zero)
	totalLogicalSectors  uint16           //  8:10 reserved (must be zero)
	MediaDescriptor      media.Descriptor // 10:11
	logicalSectorsPerFAT uint16           // 11:13 reserved (must be zero)

	// DOS 3.31 parameter block
	physicalSectorsPerTrack  uint16 // 13:15 reserved
	numberOfHeads            uint16 // 15:17 reserved
	hiddenSectors            uint32 // 17:21 reserved
	totalLogicalSectorsLarge uint32 // 21:25 reserved

	// NTFS extended parameter block
	physicalDriveNumber          byte    // 25:26
	flags                        byte    // 26:27
	extendedBootSignature        byte    // 27:28
	reserved1                    byte    // 28:29 (reserved)
	TotalSectors                 uint64  // 29:37 Number of sectors in the volume (one less than partition table entry)
	MFT                          uint64  // 37:45 Logical cluster number of $MFT
	MFTMirror                    uint64  // 45:53 Logical cluster number of $MFT mirror
	ClustersPerFileRecordSegment int8    // 53:54 if < 0, MFT record size = 2^n
	reserved2                    [3]byte // 54:57
	ClustersPerIndexBlock        int8    // 57:58
	reserved3                    [3]byte // 58:61
	VolumeSerialNumber           uint64  // 61:69
	Checksum                     uint32  // 69:73
}

// ClusterSize returns the size of an NTFS cluster in bytes.
func (block *ParameterBlock) ClusterSize() int {
	return int(block.BytesPerSector) * int(block.SectorsPerCluster)
}

// FileRecordSize returns the size of a file record segment in bytes.
func (block *ParameterBlock) FileRecordSize() int {
	if block.ClustersPerFileRecordSegment < 0 {
		return int(1) << uint(-block.ClustersPerFileRecordSegment) // 2^N
	}
	return int(block.ClustersPerFileRecordSegment) * block.ClusterSize()
}

// IndexBlockSize returns the size of an index block in bytes.
func (block *ParameterBlock) IndexBlockSize() int {
	if block.ClustersPerIndexBlock < 0 {
		return int(1) << uint(-block.ClustersPerIndexBlock) // 2^N
	}
	return int(block.ClustersPerIndexBlock) * block.ClusterSize()
}

// ReadFrom reads 73 bytes of BIOS parameter block data from r into block.
func (block *ParameterBlock) ReadFrom(r io.Reader) (n int64, err error) {
	var buf [73]byte
	n32, err := r.Read(buf[:])
	n = int64(n32)
	if err != nil {
		return
	}

	if n < ParameterBlockLength {
		return 0, ErrTruncatedData
	}

	return n, block.UnmarshalBinary(buf[:])
}

// UnmarshalBinary unmarshals the little-endian binary representation
// of BIOS parameter block data into block.
//
// The provided data must be at least 73 bytes long.
func (block *ParameterBlock) UnmarshalBinary(data []byte) error {
	if len(data) < ParameterBlockLength {
		return ErrTruncatedData
	}

	// DOS 2.0 parameter block
	block.BytesPerSector = binary.LittleEndian.Uint16(data[0:2])
	block.SectorsPerCluster = uint8(data[2])
	block.reservedSectors = binary.LittleEndian.Uint16(data[3:5])
	block.numberOfFATs = uint8(data[5])
	block.rootDirectoryEntries = binary.LittleEndian.Uint16(data[6:8])
	block.totalLogicalSectors = binary.LittleEndian.Uint16(data[8:10])
	block.MediaDescriptor = media.Descriptor(data[10])
	block.logicalSectorsPerFAT = binary.LittleEndian.Uint16(data[11:13])

	// DOS 3.31 parameter block
	block.physicalSectorsPerTrack = binary.LittleEndian.Uint16(data[13:15])
	block.numberOfHeads = binary.LittleEndian.Uint16(data[15:17])
	block.hiddenSectors = binary.LittleEndian.Uint32(data[17:21])
	block.totalLogicalSectorsLarge = binary.LittleEndian.Uint32(data[21:25])

	// NTFS extended parameter block
	block.physicalDriveNumber = data[25]
	block.flags = data[26]
	block.extendedBootSignature = data[27]
	block.reserved1 = data[28]
	block.TotalSectors = binary.LittleEndian.Uint64(data[29:37])
	block.MFT = binary.LittleEndian.Uint64(data[37:45])
	block.MFTMirror = binary.LittleEndian.Uint64(data[45:53])
	block.ClustersPerFileRecordSegment = int8(data[53])
	block.reserved2[0] = data[54]
	block.reserved2[1] = data[55]
	block.reserved2[2] = data[56]
	block.ClustersPerIndexBlock = int8(data[57])
	block.reserved3[0] = data[58]
	block.reserved3[1] = data[59]
	block.reserved3[2] = data[60]
	block.VolumeSerialNumber = binary.LittleEndian.Uint64(data[61:69])
	block.Checksum = binary.LittleEndian.Uint32(data[69:73])

	return nil
}
