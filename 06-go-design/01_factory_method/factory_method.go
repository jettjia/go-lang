package factory_method

type ReadBook interface {
	ReadBook()
}

type ReadType int

const (
	IPAD ReadType = iota + 1
	COMPUTER
)

func NewBookReader(t ReadType) ReadBook {
	switch t {
	case IPAD:
		return new(IpadFactory).CreateIpadReadEr()
	case COMPUTER:
		return new(ComputerFactory).CreateComputerReadEr()
	default:
		return nil
	}
}
