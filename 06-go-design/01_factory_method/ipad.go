package factory_method

import "fmt"

type Ipad struct{}

func (Ipad) ReadBook() {
	fmt.Println("ipad read book")
}

type IpadFactory struct {
}

func (IpadFactory) CreateIpadReadEr() ReadBook {
	return new(Ipad)
}
