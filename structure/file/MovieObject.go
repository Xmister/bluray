package file
import "io"
import "github.com/Xmister/bluray/structure"
import "encoding/binary"

type MovieObject struct {
	TypeIndicator [4]byte
	TypeIndicator2 [4]byte
	*structure.MovieObjects
}

func NewMovieObject(r io.ReadSeeker) (res *MovieObject) {
	res = &MovieObject{}
	binary.Read(r, binary.BigEndian, &res.TypeIndicator)
	binary.Read(r, binary.BigEndian, &res.TypeIndicator2)
	r.Seek(32, io.SeekCurrent)
	res.MovieObjects = structure.NewMovieObjects(r)
	return
}
