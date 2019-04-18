package structure
import "bytes"
import "io"
import "encoding/binary"

type PlayItem struct {
	OUTTime uint32
	INTime uint32
	Length uint16
	ConnectionCondition uint8
	RefToSTCID uint8
	NumberOfAngles uint8
	FileName [5]byte
	Codec [4]byte
	IsMultiAngle bool
	IsDifferentAudios bool
	IsSeamlessAudioChange bool
	Angles []*PlayItemEntry
	*STNTable
}

func NewPlayItem(r *bytes.Reader) (res *PlayItem) {
	res = &PlayItem{}
	binary.Read(r, binary.BigEndian, &res.Length)
	binary.Read(r, binary.BigEndian, &res.FileName)
	binary.Read(r, binary.BigEndian, &res.Codec)
	r.ReadByte()
	var read uint8
	binary.Read(r, binary.BigEndian, &read)
	read <<=3
	res.IsMultiAngle = (read >> 7) == 1
	res.ConnectionCondition = ((read << 1) >>4) & 0x0F
	binary.Read(r, binary.BigEndian, &res.RefToSTCID)
	binary.Read(r, binary.BigEndian, &res.INTime)
	binary.Read(r, binary.BigEndian, &res.OUTTime)
	NewUOMaskTable(r)
	r.ReadByte() //RandomAccess
	r.Seek(3, io.SeekCurrent) //Still stuff
	if res.IsMultiAngle {
		binary.Read(r, binary.BigEndian, &res.NumberOfAngles)
		binary.Read(r, binary.BigEndian, &read)
		res.IsDifferentAudios = (read & 0x01) == 1
		res.IsSeamlessAudioChange = (read & 0x02) == 1
		res.Angles = make([]*PlayItemEntry, res.NumberOfAngles)
		for i:=uint8(0); i<res.NumberOfAngles; i++ {
			res.Angles[i] = NewPlayItemEntry(r)
		}
	}
	res.STNTable = NewSTNTable(r)
	return
}
