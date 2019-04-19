package structure
import "io"
import "encoding/binary"

type StreamAttributes struct {
	Length uint8
	StreamCodingType uint8
	VideoFormat uint8
	AudioFormat uint8
	FrameRate uint8
	SampleRate uint8
	CharacterCode uint8
	LanguageCode [3]byte
}

func NewStreamAttributes(r io.ReadSeeker) (res *StreamAttributes) {
	res = &StreamAttributes{}
	binary.Read(r, binary.BigEndian, &res.Length)
	start, _ := r.Seek(0, io.SeekCurrent)
	binary.Read(r, binary.BigEndian, &res.StreamCodingType)
	var read uint8
	switch res.StreamCodingType {
	case 0x02, 0x1B, 0xEA:
		binary.Read(r, binary.BigEndian, &read)
		res.VideoFormat = read >> 4 & 0x0F
		res.FrameRate = read & 0x0F
	case 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0xA1, 0xA2:
		binary.Read(r, binary.BigEndian, &read)
		res.AudioFormat = read >> 4 & 0x0F
		res.SampleRate = read & 0x0F
		binary.Read(r, binary.BigEndian, &res.LanguageCode)
	case 0x90, 0x91:
		binary.Read(r, binary.BigEndian, &res.LanguageCode)
	case 0x92:
		binary.Read(r, binary.BigEndian, &res.CharacterCode)
		binary.Read(r, binary.BigEndian, &res.LanguageCode)
	}
	r.Seek(start+int64(res.Length), io.SeekStart)
	return
}
