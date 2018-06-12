package ntfs

import "unicode/utf16"

// utf16ToString converts data from a unicode 16 byte sequence to a string.
//
// If data is an odd number of characters an error will be returned.
func utf16ToString(data []byte) (string, error) {
	if len(data) == 0 {
		return "", nil
	}
	if len(data)%2 != 0 {
		return "", ErrInvalidUnicode
	}
	length := len(data) / 2
	buf := make([]uint16, length)
	for i := 0; i < length; i++ {
		buf[i] = uint16(data[i*2]) | uint16(data[i*2+1])<<8
	}
	return string(utf16.Decode(buf[:length])), nil
}
