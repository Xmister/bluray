package structure
import "io"
import "encoding/binary"

type Program struct {
	SPNProgramSequenceStart uint32
	ProgramMapPID           uint16
	NumberOfStreamsInPS     uint8
	Streams                 []*Stream
}

func NewProgram(r io.ReadSeeker) (res *Program) {
	res = &Program{}
	binary.Read(r, binary.BigEndian, &res.SPNProgramSequenceStart)
	binary.Read(r, binary.BigEndian, &res.ProgramMapPID)
	binary.Read(r, binary.BigEndian, &res.NumberOfStreamsInPS)
	r.Seek(1, io.SeekCurrent)
	res.Streams = make([]*Stream, res.NumberOfStreamsInPS)
	for i:=uint8(0); i<res.NumberOfStreamsInPS; i++ {
		res.Streams[i] = NewStream(r)
	}
	return
}
