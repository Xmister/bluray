package file
import "io"
import "bytes"
import "encoding/binary"
import "github.com/Xmister/bluray/structure"

type Index struct {
	TypeIndicator [4]byte
	TypeIndicator2 [4]byte
	IndexesStartAddress uint32
	ExtensionDataStartAddress uint32
	*structure.AppInfoBDMV
	*structure.Indexes
}

func NewIndex(r *bytes.Reader) (res *Index) {
	res = &Index{}
	binary.Read(r, binary.BigEndian, &res.TypeIndicator)
	binary.Read(r, binary.BigEndian, &res.TypeIndicator2)
	binary.Read(r, binary.BigEndian, &res.IndexesStartAddress)
	binary.Read(r, binary.BigEndian, &res.ExtensionDataStartAddress)
	r.Seek(24, io.SeekCurrent)
	res.AppInfoBDMV = structure.NewAppInfoBDMV(r)
	r.Seek(int64(res.IndexesStartAddress), io.SeekStart)
	res.Indexes = structure.NewIndexes(r)
	return
}
