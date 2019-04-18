package structure
import "bytes"
import "encoding/binary"

type SubPlayItem struct {
	OUTTime uint32
	SyncStartPTS uint32
	INTime uint32
	SyncPlaytItemID uint16
	Length uint16
	ConnectionCondition uint8
	RefToSTCID uint8
	NumberOfMultiClipEntries uint8
	FileName [5]byte
	Codec [4]byte
	IsMultiClipEntries bool
	MultiClipEntries []*PlayItemEntry
}

func NewSubPlayItem(r *bytes.Reader) (res *SubPlayItem) {
	res = &SubPlayItem{}
	binary.Read(r, binary.BigEndian, &res.Length)
	binary.Read(r, binary.BigEndian, &res.FileName)
	binary.Read(r, binary.BigEndian, &res.Codec)
	r.ReadByte()
	r.ReadByte()
	r.ReadByte()
	var read uint8
	binary.Read(r, binary.BigEndian, &read)
	read <<=3
	res.ConnectionCondition = read >> 4
	res.IsMultiClipEntries = ((read << 4)>>7) == 1
	binary.Read(r, binary.BigEndian, &res.RefToSTCID)
	binary.Read(r, binary.BigEndian, &res.INTime)
	binary.Read(r, binary.BigEndian, &res.OUTTime)
	binary.Read(r, binary.BigEndian, &res.SyncPlaytItemID)
	binary.Read(r, binary.BigEndian, &res.SyncStartPTS)
	if res.IsMultiClipEntries {
		binary.Read(r, binary.BigEndian, &res.NumberOfMultiClipEntries)
		r.ReadByte()
		res.MultiClipEntries = make([]*PlayItemEntry, res.NumberOfMultiClipEntries)
		for i:=uint8(0); i<res.NumberOfMultiClipEntries; i++ {
			res.MultiClipEntries[i] = NewPlayItemEntry(r)
		}

	}
	return
}
