package structure
import "io"
import "encoding/binary"

type AppInfoBDMV struct {
	Length uint32
	VideoFormat uint8
	FrameRate uint8
	UserData [8*32]byte
}

func NewAppInfoBDMV(r io.ReadSeeker) (res *AppInfoBDMV) {
	res = &AppInfoBDMV{}
	binary.Read(r, binary.BigEndian, &res.Length)
	r.Seek(1, io.SeekCurrent)
	var read uint8
	binary.Read(r, binary.BigEndian, &read)
	res.VideoFormat = read >> 4
	res.FrameRate = (read << 4) >> 4
	r.Read(res.UserData[:])
	return
}
