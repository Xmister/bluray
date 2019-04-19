package structure
import "io"

type UOMaskTable struct {
	//TODO
}

func NewUOMaskTable(r io.ReadSeeker) (res *UOMaskTable) {
	res = &UOMaskTable{}
	//Just skip the bytes for now
	r.Seek(8, io.SeekCurrent)
	return
}
