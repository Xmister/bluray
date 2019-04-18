package structure
import "bytes"
import "encoding/binary"

type PlayList struct {
	Length uint32
	NumberOfPlayItems uint16
	NumberOfSubPaths uint16
	PlayItems []*PlayItem
	SubPaths []*SubPath
}

func NewPlayList(r *bytes.Reader) (res *PlayList) {
	res = &PlayList{}
	binary.Read(r, binary.BigEndian, &res.Length)
	r.ReadByte()
	r.ReadByte()
	binary.Read(r, binary.BigEndian, &res.NumberOfPlayItems)
	binary.Read(r, binary.BigEndian, &res.NumberOfSubPaths)
	res.PlayItems = make([]*PlayItem, res.NumberOfPlayItems)
	for i:=uint16(0); i<res.NumberOfPlayItems; i++ {
		res.PlayItems[i] = NewPlayItem(r)
	}
	res.SubPaths = make([]*SubPath, res.NumberOfSubPaths)
	for i:=uint16(0); i<res.NumberOfSubPaths; i++ {
		res.SubPaths[i] = NewSubPath(r)
	}
	return
}
