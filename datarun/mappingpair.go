package datarun

// MappingPair stores the length and offset of a logical cluster number
// so that it can be mapped to a virtual cluster number.
type MappingPair struct {
	Length uint64 // In clusters
	Offset int64  // Relative to previous data run, may be negative
}
