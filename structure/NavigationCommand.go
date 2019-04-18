package structure
import "bytes"
import "encoding/binary"

type NavigationCommand struct {
	OperandCount uint8
	CommandGroup uint8
	CommandSubGroup uint8
	BranchOption uint8
	CompareOption uint8
	SetOption uint8
	ImmediateDestionation bool
	ImmediateSource bool
	Destination uint32
	Source uint32
}

func NewNavigationCommand(r *bytes.Reader) (res *NavigationCommand) {
	res = &NavigationCommand{}
	var read uint8
	binary.Read(r, binary.BigEndian, &read)
	res.OperandCount = read >> 5
	res.CommandGroup = (read << 3) >> 6
	res.CommandSubGroup = (read << 5) >> 5
	binary.Read(r, binary.BigEndian, &read)
	res.ImmediateDestionation = (read >> 7) == 1
	res.ImmediateSource = ((read << 1) >> 7) == 1
	res.BranchOption = read & 0x0F
	binary.Read(r, binary.BigEndian, &read)
	res.CompareOption = read & 0x0F
	binary.Read(r, binary.BigEndian, &read)
	res.SetOption = read & 0x1F
	binary.Read(r, binary.BigEndian, &res.Destination)
	binary.Read(r, binary.BigEndian, &res.Source)
	return
}
