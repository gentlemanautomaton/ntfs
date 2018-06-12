package ntfs

import (
	"fmt"
	"io"

	"github.com/gentlemanautomaton/ntfs/attrtype"
)

// https://en.wikipedia.org/wiki/NTFS#Master_File_Table
// http://www.kes.talktalk.net/ntfs/

// MFT provides access to the master file table of an NTFS filesystem.
type MFT struct {
	SectorSize  int64 // In bytes
	ClusterSize int64 // In bytes
	RecordSize  int64 // In bytes
	BaseAddr    int64 // In bytes
}

// File retrieves information about the file identified by id.
func (mft *MFT) File(r io.ReadSeeker, id int64) (*File, error) {
	var (
		f       File
		segment = make([]byte, mft.RecordSize)
	)

	// Read in the entire segment
	r.Seek(mft.BaseAddr+mft.RecordSize*id, io.SeekStart)
	if _, err := r.Read(segment); err != nil {
		return nil, fmt.Errorf("unable to read MFT file record data for entry %d: %v", id, err)
	}

	// Unmarshal the file record segment header
	if err := f.Header.UnmarshalBinary(segment); err != nil {
		return nil, fmt.Errorf("unable to parse file record header for entry %d: %v", id, err)
	}

	//r.Seek(int64(f.Header.FirstAttributeOffset), io.SeekCurrent)

	// Unmarshal attributes
	pos := int64(f.Header.FirstAttributeOffset)
	a := 0
	for pos < mft.RecordSize {
		data := segment[pos:]

		// Detect the end of an attribute stream
		if len(data) < 4 {
			return nil, ErrTruncatedData
		}
		if attrtype.Unmarshal(data) == attrtype.End {
			break
		}

		var attr Attribute
		if err := attr.UnmarshalBinary(data); err != nil {
			return nil, fmt.Errorf("unable to read MFT attribute %d of file record %d: %v", a, id, err)
		}

		f.Attributes = append(f.Attributes, attr)
		pos += int64(attr.Header.RecordLength)
		a++
	}

	return &f, nil
}
