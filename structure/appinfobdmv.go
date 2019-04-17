package structure
import "bytes"
import "encoding/binary"

type AppInfoBDMV struct {
	Length uint32
	VideoFormat uint8
	FrameRate uint8
	UserData [8*32]byte
}

func NewAppInfoBDMV(r *bytes.Reader) (res *AppInfoBDMV) {
	res = &AppInfoBDMV{}
	binary.Read(r, binary.BigEndian, &res.Length)
	r.ReadByte()
	binary.Read(r, binary.BigEndian, &res.VideoFormat)
	res.VideoFormat >>= 4
	r.UnreadByte()
	binary.Read(r, binary.BigEndian, &res.FrameRate)
	res.FrameRate <<= 4
	r.Read(res.UserData[:])
	return
}
