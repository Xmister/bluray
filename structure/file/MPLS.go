package file
import "io"
import "bytes"
import "encoding/binary"
import "github.com/Xmister/bluray/structure"

type MPLS struct {
	TypeIndicator            [4]byte
	TypeIndicator2           [4]byte
	PlayListStartAddress     uint32
	PlayListMarkStartAddress uint32
	ExtensionDataStartAddress uint32
	*structure.AppInfoPlayList
	*structure.PlayList
}

func NewMPLS(r *bytes.Reader) (res *MPLS) {
	res = &MPLS{}
	binary.Read(r, binary.BigEndian, &res.TypeIndicator)
	binary.Read(r, binary.BigEndian, &res.TypeIndicator2)
	binary.Read(r, binary.BigEndian, &res.PlayListStartAddress)
	binary.Read(r, binary.BigEndian, &res.PlayListMarkStartAddress)
	binary.Read(r, binary.BigEndian, &res.ExtensionDataStartAddress)
	r.Seek(20, io.SeekCurrent)
	res.AppInfoPlayList = structure.NewAppInfoPlayList(r)
	r.Seek(int64(res.PlayListStartAddress), io.SeekStart)
	res.PlayList = structure.NewPlayList(r)
	return
}
