package attrtype

import (
	"encoding/binary"
	"strconv"
)

// Code is an attribute type code.
type Code uint32

// Attribute type codes.
const (
	StandardInformation Code = 0x0010     // $STANDARD_INFORMATION
	AttributeList       Code = 0x0020     // $ATTRIBUTE_LIST
	FileName            Code = 0x0030     // $FILE_NAME
	ObjectID            Code = 0x0040     // $OBJECT_ID
	SecurityDescriptor  Code = 0x0050     // $SECURITY_DESCRIPTOR
	VolumeName          Code = 0x0060     // $VOLUME_NAME
	VolumeInformation   Code = 0x0070     // $VOLUME_INFORMATION
	Data                Code = 0x0080     // $DATA
	IndexRoot           Code = 0x0090     // $INDEX_ROOT
	IndexAllocation     Code = 0x00A0     // $INDEX_ALLOCATION
	Bitmap              Code = 0x00B0     // $BITMAP
	ReparsePoint        Code = 0x00C0     // $REPARSE_POINT aka $SYMBOLIC_LINK
	EAInformation       Code = 0x00D0     // $EA_INFORMATION
	EA                  Code = 0x00E0     // $EA
	PropertySet         Code = 0x00F0     // $PROPERTY_SET
	UserDefined         Code = 0x0100     // $FIRST_USER_DEFINED_ATTRIBUTE
	End                 Code = 0xFFFFFFFF // End of attribute stream
)

// String returns a description of the attribute type code.
func (c Code) String() string {
	switch c {
	case StandardInformation:
		return "$STANDARD_INFORMATION"
	case AttributeList:
		return "$ATTRIBUTE_LIST"
	case FileName:
		return "$FILE_NAME"
	case ObjectID:
		return "$OBJECT_ID"
	case SecurityDescriptor:
		return "$SECURITY_DESCRIPTOR"
	case VolumeName:
		return "$VOLUME_NAME"
	case VolumeInformation:
		return "$VOLUME_INFORMATION"
	case Data:
		return "$DATA"
	case IndexRoot:
		return "$INDEX_ROOT"
	case IndexAllocation:
		return "$INDEX_ALLOCATION"
	case Bitmap:
		return "$BITMAP"
	case ReparsePoint:
		return "$REPARSE_POINT"
	case EAInformation:
		return "$EA_INFORMATION"
	case EA:
		return "$EA"
	case PropertySet:
		return "$PROPERTY_SET"
	case End:
		return "END"
	default:
		return "USER_DEFINED(" + strconv.Itoa(int(c)) + ")"
	}
}

// Unmarshal unmarshals the little-endian binary representation
// of an attribute type code.
//
// The provided data must be at least 4 bytes long, or unmarshal will
// panic.
func Unmarshal(data []byte) Code {
	return Code(binary.LittleEndian.Uint32(data[0:4]))
}
