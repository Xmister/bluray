package bluray
import "bytes"
import "io"
import "io/ioutil"
import "github.com/Xmister/bluray/structure"

type BluRay struct {

}

func (b *BluRay) ReadIndex(r io.Reader) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	structure.NewIndx(bytes.NewReader(buf))
}