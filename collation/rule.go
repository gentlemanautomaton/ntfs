package collation

import (
	"encoding/binary"
	"strconv"
)

// Rule is an ntfs collation rule.
type Rule uint32

// NTFS collation rules.
const (
	Binary   Rule = iota // COLLATION_BINARY
	FileName             // COLLATION_FILE_NAME
	Unicode              // COLLATION_UNICODE_STRING
)

// String returns a description of the ntfs collation rule.
func (r Rule) String() string {
	switch r {
	case Binary:
		return "BINARY"
	case FileName:
		return "FILENAME"
	case Unicode:
		return "UNICODE"
	default:
		return "COLLATION(" + strconv.Itoa(int(r)) + ")"
	}
}

// Unmarshal unmarshals the little-endian binary representation
// of an ntfs collation rule.
//
// The provided data must be at least 4 bytes long, or unmarshal will
// panic.
func Unmarshal(data []byte) Rule {
	return Rule(binary.LittleEndian.Uint32(data[0:4]))
}
