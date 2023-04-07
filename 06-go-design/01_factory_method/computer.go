package factory_method

import "fmt"

type Computer struct{}

func (Computer) ReadBook() {
	fmt.Println("computer read book")
}

type ComputerFactory struct {
}

func (ComputerFactory) CreateComputerReadEr() ReadBook {
	return new(Computer)
}
