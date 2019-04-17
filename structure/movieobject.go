package structure
import "io"
import "bytes"
import "encoding/binary"

type MovieObject struct {
	ResumeIntentionFlag uint8
	MenuCallMask uint8
	TitleSearchMask uint8
	NumberOfNavigationCommands uint16
	NavigationCommands []*NavigationCommand
}

func NewMovieObject(r *bytes.Reader) (res *MovieObject) {
	res = &MovieObject{}
	var read uint8
	binary.Read(r, binary.BigEndian, &read)
	res.ResumeIntentionFlag = read & 2^0
	res.MenuCallMask = read & 2^1
	res.TitleSearchMask = read & 2^2
	r.ReadByte()
	binary.Read(r, binary.BigEndian, &res.NumberOfNavigationCommands)
	res.NavigationCommands = make([]*NavigationCommand, res.NumberOfNavigationCommands)
	for i:=uint16(0); i<res.NumberOfNavigationCommands; i++ {
		res.NavigationCommands[i] = NewNavigationCommand(r)
	}
	return
}
