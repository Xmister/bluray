package structure
import "io"
import "encoding/binary"

type MovieObjects struct {
	Length uint32
	NumberOfMovieObjects uint16
	MovieObjects []*MovieObject
}

func NewMovieObjects(r io.ReadSeeker) (res *MovieObjects) {
	res = &MovieObjects{}
	binary.Read(r, binary.BigEndian, &res.Length)
	r.Seek(4, io.SeekCurrent)
	binary.Read(r, binary.BigEndian, &res.NumberOfMovieObjects)
	res.MovieObjects = make([]*MovieObject, res.NumberOfMovieObjects)
	for i:=uint16(0); i<res.NumberOfMovieObjects; i++ {
		res.MovieObjects[i] = NewMovieObject(r)
	}
	return
}
