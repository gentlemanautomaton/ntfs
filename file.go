package ntfs

// File represents a file within an NTFS master file table.
type File struct {
	Header     FileRecordSegmentHeader
	Attributes []Attribute
}
