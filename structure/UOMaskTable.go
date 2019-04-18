package structure
import "bytes"
import "io"

type UOMaskTable struct {
	//TODO
}

func NewUOMaskTable(r *bytes.Reader) (res *UOMaskTable) {
	res = &UOMaskTable{}
	//Just skip the bytes for now
	r.Seek(8, io.SeekCurrent)
	return
}
