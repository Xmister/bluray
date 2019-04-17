package structure
import "io"
import "bytes"
import "encoding/binary"

type NavigationCommand struct {


}

func NewNavigationCommand(r *bytes.Reader) (res *NavigationCommand) {
	res = &NavigationCommand{}
	binary.Read(r, binary.BigEndian, &res.TypeIndicator)
	return
}
