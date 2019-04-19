package structure
import "io"
import "encoding/binary"

type PlayList struct {
	Length uint32
	NumberOfPlayItems uint16
	NumberOfSubPaths uint16
	PlayItems []*PlayItem
	SubPaths []*SubPath
}

func NewPlayList(r io.ReadSeeker) (res *PlayList) {
	res = &PlayList{}
	binary.Read(r, binary.BigEndian, &res.Length)
	r.Seek(1, io.SeekCurrent)
	r.Seek(1, io.SeekCurrent)
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
