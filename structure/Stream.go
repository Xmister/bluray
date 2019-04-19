package structure
import "io"
import "encoding/binary"

type Stream struct {
	StreamPID           uint16
	*StreamAttributes
}

func NewStream(r io.ReadSeeker) (res *Stream) {
	res = &Stream{}
	binary.Read(r, binary.BigEndian, &res.StreamPID)
	res.StreamAttributes = NewStreamAttributes(r)
	return
}
