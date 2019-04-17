package structure
import "bytes"
import "encoding/binary"

type Indexes struct {
	Length uint32
	FirstPlaybackTitle *Title
	TopMenuTitle *Title
	NumberOfTitles uint16
	Titles []*Title
}

func NewIndexes(r *bytes.Reader) (res *Indexes) {
	res = &Indexes{}
	binary.Read(r, binary.BigEndian, &res.Length)
	res.FirstPlaybackTitle = NewTitle(r)
	res.TopMenuTitle = NewTitle(r)
	binary.Read(r, binary.BigEndian, &res.NumberOfTitles)
	res.Titles = make([]*Title, res.NumberOfTitles)
	for i:=uint16(0); i<res.NumberOfTitles; i++ {
		res.Titles[i] = NewTitle(r)
	}
	return
}
