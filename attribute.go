package ntfs

import (
	"github.com/gentlemanautomaton/ntfs/attrtype"
)

// Attribute holds file attribute data.
type Attribute struct {
	Header        AttributeRecordHeader
	Resident      ResidentAttributeRecordHeader    // When in resident form
	Nonresident   NonresidentAttributeRecordHeader // When in non-resident form
	Name          string
	ResidentValue []byte
}

// ResidentValueString returns the value of resident attributes as a string.
//
// If the attribute is non-resident an empty string will be returned.
func (attr *Attribute) ResidentValueString() (string, error) {
	if !attr.Header.Resident() {
		return "", nil
	}
	switch attr.Header.TypeCode {
	case attrtype.StandardInformation:
		var si StandardInformation
		err := si.UnmarshalBinary(attr.ResidentValue)
		return si.String(), err
	case attrtype.FileName:
		var fn FileName
		err := fn.UnmarshalBinary(attr.ResidentValue)
		return fn.Value, err
	case attrtype.ObjectID:
		var oid ObjectID
		err := oid.UnmarshalBinary(attr.ResidentValue)
		return oid.String(), err
	case attrtype.VolumeInformation:
		var vi VolumeInformation
		err := vi.UnmarshalBinary(attr.ResidentValue)
		return vi.String(), err
	case attrtype.VolumeName:
		return utf16ToString(attr.ResidentValue)
	default:
		return "", nil
	}
}

// UnmarshalBinary unmarshals the little-endian binary representation
// of an attribute record into attr.
func (attr *Attribute) UnmarshalBinary(data []byte) error {
	// Read the common portion of the attribute record header
	if err := attr.Header.UnmarshalBinary(data); err != nil {
		return err
	}

	// Sanity check the record length
	if int(attr.Header.RecordLength) > len(data) {
		return ErrTruncatedData
	}

	// For the sake of simple bounds checking below, restrict our working
	// data set to this record
	if len(data) > int(attr.Header.RecordLength) {
		data = data[:attr.Header.RecordLength]
	}

	// Read the form-specific portion of the attribute record header
	formHeader := data[AttributeRecordHeaderLength:]
	if attr.Header.Resident() {
		if err := attr.Resident.UnmarshalBinary(formHeader); err != nil {
			return err
		}
	} else {
		if err := attr.Nonresident.UnmarshalBinary(formHeader); err != nil {
			return err
		}
	}

	// Read the attribute name if it has one
	if attr.Header.NameLength > 0 {
		start := int(attr.Header.NameOffset)
		length := int(attr.Header.NameLength) * 2 // NameLength is in characters, assuming 16-bit unicode
		end := start + length
		if end > len(data) {
			return ErrAttributeNameOutOfBounds
		}
		var err error
		attr.Name, err = utf16ToString(data[start:end])
		if err != nil {
			return err
		}
	}

	// Read the attribute value if it's resident and nonzero
	if attr.Header.Resident() && attr.Resident.ValueLength > 0 {
		start := int(attr.Resident.ValueOffset)
		length := int(attr.Resident.ValueLength)
		end := start + length
		if end > len(data) {
			return ErrAttributeValueOutOfBounds
		}
		attr.ResidentValue = make([]byte, length)
		copy(attr.ResidentValue, data[start:end])
	}

	return nil
}
