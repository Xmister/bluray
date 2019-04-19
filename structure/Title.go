package structure
import "io"
import "encoding/binary"

type Title struct {
	ObjectType uint8
	AccesType uint8
	PlaybackType uint8
	RefToMovieObjectID uint16
	RefToBDJObjectID uint32
}

func NewTitle(r io.ReadSeeker) (res *Title) {
	res = &Title{}
	var read uint8
	binary.Read(r, binary.BigEndian, &read)
	res.ObjectType = read >> 6
	res.AccesType = (read << 2) >> 6
	r.Seek(3, io.SeekCurrent)
	binary.Read(r, binary.BigEndian, &res.PlaybackType)
	res.PlaybackType >>= 6
	r.Seek(1, io.SeekCurrent)
	if res.ObjectType == 1 {
		binary.Read(r, binary.BigEndian, &res.RefToMovieObjectID)
		r.Seek(4, io.SeekCurrent)
	} else {
		binary.Read(r, binary.BigEndian, &res.RefToBDJObjectID)
		r.Seek(2, io.SeekCurrent)
	}
	return
}
