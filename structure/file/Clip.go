package file
import "io"
import "encoding/binary"
import "github.com/Xmister/bluray/structure"

type Clip struct {
	TypeIndicator [4]byte
	TypeIndicator2 [4]byte
	SequenceInfoStartAddress uint32
	ProgramInfoStartAddress uint32
	*structure.ProgramInfo
}

func NewClip(r io.ReadSeeker) (res *Clip) {
	res = &Clip{}
	binary.Read(r, binary.BigEndian, &res.TypeIndicator)
	binary.Read(r, binary.BigEndian, &res.TypeIndicator2)
	binary.Read(r, binary.BigEndian, &res.SequenceInfoStartAddress)
	binary.Read(r, binary.BigEndian, &res.ProgramInfoStartAddress)
	r.Seek(int64(res.ProgramInfoStartAddress), io.SeekStart)
	res.ProgramInfo = structure.NewProgramInfo(r)
	return
}
