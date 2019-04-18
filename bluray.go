package bluray
import "bytes"
import "io"
import "io/ioutil"
import "github.com/Xmister/bluray/structure/file"

type BluRay struct {
	*file.Index
	*file.MovieObject
	*file.MPLS
}

func (b *BluRay) ReadIndex(r io.Reader) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	b.Index = file.NewIndex(bytes.NewReader(buf))
}

func (b *BluRay) ReadMovieObject(r io.Reader) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	b.MovieObject = file.NewMovieObject(bytes.NewReader(buf))
}

func (b *BluRay) ReadMPLS(r io.Reader) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	b.MPLS = file.NewMPLS(bytes.NewReader(buf))
}