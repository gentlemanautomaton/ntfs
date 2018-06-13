package ntfs

import "github.com/google/uuid"

// GUID is a 16-byte array.
type GUID = uuid.UUID

func unmarshalGUID(data []byte) GUID {
	var g GUID
	if err := g.UnmarshalBinary(data); err != nil {
		return g
	}
	// Byte-order conversion from little-endian to big-endian
	g[0], g[1], g[2], g[3] = g[3], g[2], g[1], g[0]
	g[4], g[5] = g[5], g[4]
	g[6], g[7] = g[7], g[6]
	return g
}
