package structure
import "io"
import "encoding/binary"

type stream struct {
	*StreamEntry
	*StreamAttributes
}

type STNTable struct {
	Length uint16
	NumberOfPrimaryVideoStreamEntries uint8
	NumberOfPrimaryAudioStreamEntries uint8
	NumberOfPrimaryPGStreamEntries uint8
	NumberOfPrimaryIGStreamEntries uint8
	NumberOfSecondaryAudioStreamEntries uint8
	NumberOfSecondaryVideoStreamEntries uint8
	NumberOfSecondaryPGStreamEntries uint8
	PrimaryVideoStreamEntries []stream
	PrimaryAudioStreamEntries []stream
	PrimaryPGStreamEntries []stream
	PrimaryIGStreamEntries []stream
	SecondaryAudioStreamEntries []stream
	SecondaryVideoStreamEntries []stream
	SecondaryPGStreamEntries []stream
}

func loadStreams(r io.ReadSeeker, n uint8) (res []stream) {
	res = make([]stream, n)
	for i := uint8(0); i<n; i++ {
		res[i].StreamEntry = NewStreamEntry(r)
		res[i].StreamAttributes = NewStreamAttributes(r)
	}
	return
}

func NewSTNTable(r io.ReadSeeker) (res *STNTable) {
	res = &STNTable{}
	binary.Read(r, binary.BigEndian, &res.Length)
	r.Seek(1, io.SeekCurrent)
	r.Seek(1, io.SeekCurrent)
	binary.Read(r, binary.BigEndian, &res.NumberOfPrimaryVideoStreamEntries)
	binary.Read(r, binary.BigEndian, &res.NumberOfPrimaryAudioStreamEntries)
	binary.Read(r, binary.BigEndian, &res.NumberOfPrimaryPGStreamEntries)
	binary.Read(r, binary.BigEndian, &res.NumberOfPrimaryIGStreamEntries)
	binary.Read(r, binary.BigEndian, &res.NumberOfSecondaryAudioStreamEntries)
	binary.Read(r, binary.BigEndian, &res.NumberOfSecondaryVideoStreamEntries)
	binary.Read(r, binary.BigEndian, &res.NumberOfSecondaryPGStreamEntries)
	r.Seek(5, io.SeekCurrent)
	res.PrimaryVideoStreamEntries = loadStreams(r, res.NumberOfPrimaryVideoStreamEntries)
	res.PrimaryAudioStreamEntries = loadStreams(r, res.NumberOfPrimaryAudioStreamEntries)
	res.PrimaryPGStreamEntries = loadStreams(r, res.NumberOfPrimaryPGStreamEntries)
	res.SecondaryPGStreamEntries = loadStreams(r, res.NumberOfSecondaryPGStreamEntries)
	res.PrimaryIGStreamEntries = loadStreams(r, res.NumberOfPrimaryIGStreamEntries)
	res.SecondaryAudioStreamEntries = loadStreams(r, res.NumberOfSecondaryAudioStreamEntries)
	res.SecondaryVideoStreamEntries = loadStreams(r, res.NumberOfSecondaryVideoStreamEntries)
	return
}
