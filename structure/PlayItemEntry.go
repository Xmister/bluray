package structure
import "io"
import "encoding/binary"

type PlayItemEntry struct {
	RefToSTCID uint8
	FileName [5]byte
	Codec [4]byte
}

func NewPlayItemEntry(r io.ReadSeeker) (res *PlayItemEntry) {
	res = &PlayItemEntry{}
	binary.Read(r, binary.BigEndian, &res.FileName)
	binary.Read(r, binary.BigEndian, &res.Codec)
	binary.Read(r, binary.BigEndian, &res.RefToSTCID)
	return
}
