package ntfs

// ObjectIDMinLength is the minimum length of an object ID attribute in bytes.
const ObjectIDMinLength = 16

// ObjectID holds object ID attribute data.
type ObjectID struct {
	Value         GUID
	BirthVolumeID GUID
	BirthObjectID GUID
	DomainID      GUID
}

// UnmarshalBinary unmarshals the little-endian binary representation
// of an object ID attribute into id.
//
// The provided data must be at least 16 bytes long.
func (id *ObjectID) UnmarshalBinary(data []byte) error {
	if len(data) < ObjectIDMinLength {
		return ErrTruncatedData
	}
	id.Value = unmarshalGUID(data[0:16])
	if len(data) >= 32 {
		id.BirthVolumeID = unmarshalGUID(data[16:32])
	}
	if len(data) >= 48 {
		id.BirthObjectID = unmarshalGUID(data[32:48])
	}
	if len(data) >= 64 {
		id.DomainID = unmarshalGUID(data[48:64])
	}
	return nil
}

// String returns a description of the object ID.
func (id *ObjectID) String() string {
	return id.Value.String()
}
