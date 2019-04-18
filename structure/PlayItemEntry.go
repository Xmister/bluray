package structure
import "bytes"
import "encoding/binary"

type PlayItemEntry struct {
	RefToSTCID uint8
	FileName [5]byte
	Codec [4]byte
}

func NewPlayItemEntry(r *bytes.Reader) (res *PlayItemEntry) {
	res = &PlayItemEntry{}
	binary.Read(r, binary.BigEndian, &res.FileName)
	binary.Read(r, binary.BigEndian, &res.Codec)
	binary.Read(r, binary.BigEndian, &res.RefToSTCID)
	return
}
