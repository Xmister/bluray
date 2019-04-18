package structure
import "bytes"
import "encoding/binary"

type StreamEntry struct {
	Length uint8
	StreamType uint8
	RefToStreamPID uint16
	RefToSubPathID uint8
	RefToSubClipID uint8
}

func NewStreamEntry(r *bytes.Reader) (res *StreamEntry) {
	res = &StreamEntry{}
	binary.Read(r, binary.BigEndian, &res.Length)
	if res.Length == 0 {
		return
	}
	binary.Read(r, binary.BigEndian, &res.StreamType)
	switch res.StreamType {
	case 1, 3:
		binary.Read(r, binary.BigEndian, &res.RefToStreamPID)
	case 2, 4:
		binary.Read(r, binary.BigEndian, &res.RefToSubPathID)
		binary.Read(r, binary.BigEndian, &res.RefToSubClipID)
		binary.Read(r, binary.BigEndian, &res.RefToStreamPID)
	}
	return
}
