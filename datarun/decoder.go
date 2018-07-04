package datarun

import "io"

// https://msdn.microsoft.com/library/bb470039

// Decoder decodes data in an underlying reader as a serires of data run
// mapping pairs.
type Decoder struct {
	r      io.Reader
	offset int64
	done   bool
}

// NewDecoder returns a new decoder that reads from r until it encounters an
// end-of-stream header.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

// Read returns the next mapping pair from d. It returns io.EOF when there is no
// more data to return.
func (d *Decoder) Read() (MappingPair, error) {
	// If we've already encountered the end-of-stream sentinal value stop now.
	if d.done {
		return MappingPair{}, io.EOF
	}

	// Read in the header (1 byte)
	var buf [16]byte
	if err := fillBuffer(d.r, buf[0:1]); err != nil {
		return MappingPair{}, err
	}

	// Make sure the header is valid
	header := Header(buf[0])
	if !header.Valid() {
		return MappingPair{}, ErrInvalidHeader
	}

	// Check whether this header indicates the end of the stream
	if header == EOF {
		d.done = true
		return MappingPair{}, io.EOF
	}

	// Read in the offset and length values (1-16 bytes)
	if err := fillBuffer(d.r, buf[0:header.DataBytes()]); err != nil {
		return MappingPair{}, err
	}

	var result MappingPair

	// Parse the length
	lenBytes := header.LengthBytes()
	if lenBytes > 0 {
		result.Length = parseUint64(buf[0:lenBytes])
	}

	// Parse the offset and add it to the previous value
	offBytes := header.OffsetBytes()
	if offBytes > 0 {
		off1 := lenBytes
		off2 := off1 + header.OffsetBytes()
		d.offset += parseInt64(buf[off1:off2])
		result.Offset = d.offset
	}

	return result, nil
}

func fillBuffer(r io.Reader, b []byte) error {
	n, err := r.Read(b)
	if n != len(b) {
		if err == io.EOF || err == nil {
			return ErrTruncatedData
		}
		return err
	}
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}
