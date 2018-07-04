package datarun

func parseUint64(data []byte) (value uint64) {
	for i, k := uint(0), uint(len(data)); i < k; i++ {
		value |= uint64(data[i]) << i * 8
	}
	return
}

func parseInt64(data []byte) (value int64) {
	// FIXME: Preserve sign?
	for i, k := uint(0), uint(len(data)); i < k; i++ {
		value |= int64(data[i]) << i * 8
	}
	return
}
