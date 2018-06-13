package ntfs

import (
	"encoding/binary"
	"time"
)

const windowsTimeFormat = "2006-01-02 15:04:05 MST"

// Windows Epoch: 1601-01-01 00:00:00 UTC
//    Unix Epoch: 1970-01-01 00:00:00 UTC
//    Difference: 11644473600 seconds

const (
	// The offset between the windows epoch and unix epoch in 1-second intervals
	unixtimeOffsetSeconds = 11644473600

	// The offset between the windows epoch and unix epoch in 100-nanosecond intervals
	unixtimeOffset100Nano = unixtimeOffsetSeconds * (int64(time.Second) / 100)
)

func unmarshalFileTime(data []byte) time.Time {
	t := int64(binary.LittleEndian.Uint64(data)) // 100 nanosecond intervals since the windows epoch
	t -= unixtimeOffset100Nano                   // converted to unixtime (in 100 nanoseconds)
	t *= 100                                     // converted to unixtime (in nanoseconds)
	return time.Unix(0, t).UTC()
}

func localTimeString(t time.Time) string {
	return t.Local().Format(windowsTimeFormat)
}
