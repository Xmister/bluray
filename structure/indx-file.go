package structure
import "io"
import "bytes"
import "encoding/binary"

type Indx struct {
	TypeIndicator [4]byte
	TypeIndicator2 [4]byte
	IndexesStartAddress uint32
	ExtensionDataStartAddress uint32
	*AppInfoBDMV
	*Indexes
}

func NewIndx(r *bytes.Reader) (res *Indx) {
	res = &Indx{}
	binary.Read(r, binary.BigEndian, &res.TypeIndicator)
	binary.Read(r, binary.BigEndian, &res.TypeIndicator2)
	binary.Read(r, binary.BigEndian, &res.IndexesStartAddress)
	binary.Read(r, binary.BigEndian, &res.ExtensionDataStartAddress)
	r.Seek(24, io.SeekCurrent)
	res.AppInfoBDMV = NewAppInfoBDMV(r)
	r.Seek(int64(res.IndexesStartAddress), io.SeekStart)
	res.Indexes = NewIndexes(r)
	return
}
