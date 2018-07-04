package datarun

import (
	"bytes"
	"io"
)

// Decode decodes data until it encounters EOF in the data stream or encounters an
// error. It returns a slice of all the mapping pairs it decoded. If the data is
// decoded successfully err will be nil.
func Decode(data []byte) (pairs []MappingPair, err error) {
	d := NewDecoder(bytes.NewReader(data))
	pair, err := d.Read()
	for err == nil {
		pairs = append(pairs, pair)
		pair, err = d.Read()
	}
	if err == io.EOF {
		err = nil
	}
	return
}
