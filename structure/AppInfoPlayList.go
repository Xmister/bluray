package structure
import "io"
import "encoding/binary"

type AppInfoPlayList struct {
	Length uint32
	PlaybackType uint8
	PlaybackCount uint16
}

func NewAppInfoPlayList(r io.ReadSeeker) (res *AppInfoPlayList) {
	res = &AppInfoPlayList{}
	binary.Read(r, binary.BigEndian, &res.Length)
	r.Seek(1, io.SeekCurrent)
	binary.Read(r, binary.BigEndian, &res.PlaybackType)
	if res.PlaybackType >= 2 && res.PlaybackType <=3 {
		binary.Read(r, binary.BigEndian, &res.PlaybackCount)
	}
	return
}
