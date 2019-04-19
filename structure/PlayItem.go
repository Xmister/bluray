package structure
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

func NewPlayItem(r io.ReadSeeker) (res *PlayItem) {
	res = &PlayItem{}
	binary.Read(r, binary.BigEndian, &res.Length)
	binary.Read(r, binary.BigEndian, &res.FileName)
	binary.Read(r, binary.BigEndian, &res.Codec)
	r.Seek(1, io.SeekCurrent)
	var read uint8
	binary.Read(r, binary.BigEndian, &read)
	read <<=3
	res.IsMultiAngle = (read >> 7) == 1
	res.ConnectionCondition = ((read << 1) >>4) & 0x0F
	binary.Read(r, binary.BigEndian, &res.RefToSTCID)
	binary.Read(r, binary.BigEndian, &res.INTime)
	binary.Read(r, binary.BigEndian, &res.OUTTime)
	NewUOMaskTable(r)
	r.Seek(1, io.SeekCurrent) //RandomAccess
	r.Seek(3, io.SeekCurrent) //Still stuff
	if res.IsMultiAngle {
		binary.Read(r, binary.BigEndian, &res.NumberOfAngles)
		binary.Read(r, binary.BigEndian, &read)
		res.IsDifferentAudios = (read & 0x01) == 1
		res.IsSeamlessAudioChange = (read & 0x02) == 1
		res.Angles = make([]*PlayItemEntry, res.NumberOfAngles)
		res.Angles[0] = &PlayItemEntry{
			RefToSTCID: res.RefToSTCID,
			FileName: res.FileName,
			Codec: res.Codec,
		}
		for i:=uint8(1); i<res.NumberOfAngles; i++ {
			res.Angles[i] = NewPlayItemEntry(r)
		}
	} else {
		res.Angles = []*PlayItemEntry{
			&PlayItemEntry{
				RefToSTCID: res.RefToSTCID,
				FileName: res.FileName,
				Codec: res.Codec,
			},
		}
	}
	res.STNTable = NewSTNTable(r)
	return
}
