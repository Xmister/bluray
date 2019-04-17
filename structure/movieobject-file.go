package structure
import "io"
import "bytes"
import "encoding/binary"

type MovieObjectFile struct {
	TypeIndicator [4]byte
	TypeIndicator2 [4]byte
	*MovieObjects
}

func NewMovieObjectFile(r *bytes.Reader) (res *MovieObjectFile) {
	res = &MovieObjectFile{}
	binary.Read(r, binary.BigEndian, &res.TypeIndicator)
	binary.Read(r, binary.BigEndian, &res.TypeIndicator2)
	r.Seek(32, io.SeekCurrent)
	res.MovieObjects = NewMovieObjects(r)
	return
}
