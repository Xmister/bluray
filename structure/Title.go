package structure
import "bytes"
import "io"
import "encoding/binary"

type Title struct {
	ObjectType uint8
	AccesType uint8
	PlaybackType uint8
	RefToMovieObjectID uint16
	RefToBDJObjectID uint32
}

func NewTitle(r *bytes.Reader) (res *Title) {
	res = &Title{}
	binary.Read(r, binary.BigEndian, &res.ObjectType)
	res.ObjectType >>= 6
	r.UnreadByte()
	binary.Read(r, binary.BigEndian, &res.AccesType)
	res.AccesType <<= 2
	res.AccesType >>= 6
	r.Seek(3, io.SeekCurrent)
	binary.Read(r, binary.BigEndian, &res.PlaybackType)
	res.PlaybackType >>= 6
	r.ReadByte()
	if res.ObjectType == 1 {
		binary.Read(r, binary.BigEndian, &res.RefToMovieObjectID)
		r.Seek(4, io.SeekCurrent)
	} else {
		binary.Read(r, binary.BigEndian, &res.RefToBDJObjectID)
		r.Seek(2, io.SeekCurrent)
	}
	return
}
