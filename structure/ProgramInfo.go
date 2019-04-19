package structure
import "io"
import "encoding/binary"

type ProgramInfo struct {
	Length           uint32
	NumberOfPrograms uint8
	Programs         []*Program
}

func NewProgramInfo(r io.ReadSeeker) (res *ProgramInfo) {
	res = &ProgramInfo{}
	binary.Read(r, binary.BigEndian, &res.Length)
	r.Seek(1, io.SeekCurrent)
	binary.Read(r, binary.BigEndian, &res.NumberOfPrograms)
	res.Programs = make([]*Program, res.NumberOfPrograms)
	for i:=uint8(0); i<res.NumberOfPrograms; i++ {
		res.Programs[i] = NewProgram(r)
	}
	return
}
