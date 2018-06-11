package attrform

// Code is an attribute form code.
type Code uint8

// Attribute form codes.
const (
	Resident    Code = 0x00 // RESIDENT_FORM
	Nonresident Code = 0x01 // NONRESIDENT_FORM
)

// String returns a description of the attribute form code.
func (c Code) String() string {
	switch c {
	case Resident:
		return "RESIDENT"
	case Nonresident:
		return "NONRESIDENT"
	default:
		return "UNKNOWN"
	}
}
