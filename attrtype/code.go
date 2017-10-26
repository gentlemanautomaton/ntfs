package attrtype

// Code is an attribute type code.
//
// TODO: Verify that uint32 is the correct underlying type.
type Code uint32

// Attribute record type codes.
const (
	StandardInformation Code = 0x10 // $STANDARD_INFORMATION
	AttributeList       Code = 0x20 // $ATTRIBUTE_LIST
	FileName            Code = 0x30 // $FILE_NAME
	ObjectID            Code = 0x40 // $OBJECT_ID
	VolumeName          Code = 0x60 // $VOLUME_NAME
	VolumeInformation   Code = 0x70 // $VOLUME_INFORMATION
	Data                Code = 0x80 // $DATA
	IndexRoot           Code = 0x90 // $INDEX_ROOT
	IndexAllocation     Code = 0xA0 // $INDEX_ALLOCATION
	Bitmap              Code = 0xB0 // $BITMAP
	ReparsePoint        Code = 0xC0 // $REPARSE_POINT
)
