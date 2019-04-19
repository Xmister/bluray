package structure
import "io"
import "encoding/binary"

type SubPath struct {
	Length uint32
	SubPathType uint8
	NumberOfSubPlayItems uint8
	IsRepeatSubPath bool
	SubPlayItems []*SubPlayItem
}

func NewSubPath(r io.ReadSeeker) (res *SubPath) {
	res = &SubPath{}
	binary.Read(r, binary.BigEndian, &res.Length)
	r.Seek(1, io.SeekCurrent)
	binary.Read(r, binary.BigEndian, &res.SubPathType)
	r.Seek(1, io.SeekCurrent)
	var read uint8
	binary.Read(r, binary.BigEndian, &read)
	res.IsRepeatSubPath = (read & 0x01) == 1
	binary.Read(r, binary.BigEndian, &res.NumberOfSubPlayItems)
	res.SubPlayItems = make([]*SubPlayItem, res.NumberOfSubPlayItems)
	for i:=uint8(0); i<res.NumberOfSubPlayItems; i++ {
		res.SubPlayItems[i] = NewSubPlayItem(r)
	}
	return
}
