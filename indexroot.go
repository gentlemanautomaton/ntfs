package ntfs

import (
	"encoding/binary"
	"fmt"

	"github.com/gentlemanautomaton/ntfs/attrtype"
	"github.com/gentlemanautomaton/ntfs/collation"
)

// IndexRootMinLength is the minimum length of an $INDEX_ROOT attribute
// in bytes.
const IndexRootMinLength = 16

// IndexRoot stores $INDEX_ROOT attribute information.
type IndexRoot struct {
	AttrType             attrtype.Code  //  0:4 The type of attributes being indexed
	CollationRule        collation.Rule //  4:8
	BytesPerIndexRecord  uint32         //  8:12
	BlocksPerIndexRecord uint8          // 12:13 In sectors if BPIR < ClusterSize, otherwise clusters
	reserved1            [3]byte        // 13:15
}

// UnmarshalBinary unmarshals the little-endian binary representation
// of an $INDEX_ROOT attribute into index.
func (index *IndexRoot) UnmarshalBinary(data []byte) error {
	if len(data) < IndexRootMinLength {
		return ErrTruncatedData
	}
	index.AttrType = attrtype.Unmarshal(data[0:4])
	index.CollationRule = collation.Unmarshal(data[4:8])
	index.BytesPerIndexRecord = binary.LittleEndian.Uint32(data[8:12])
	index.BlocksPerIndexRecord = data[12]
	index.reserved1[0] = data[13]
	index.reserved1[1] = data[14]
	index.reserved1[2] = data[15]
	return nil
}

// String returns a description of the volume information.
func (index *IndexRoot) String() string {
	return fmt.Sprintf("AttrType: %s, Collation: %s, IndexRecordSize: %d bytes, Blocks: %d",
		index.AttrType, index.CollationRule, index.BytesPerIndexRecord, index.BlocksPerIndexRecord)
}
