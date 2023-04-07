package adapter

import "fmt"

// 国外中锋
type ForeignCenter struct {
	name string
}

func NewForeignCenter(name string) *ForeignCenter {
	return &ForeignCenter{
		name: name,
	}
}

func (c ForeignCenter) Name() string {
	return c.name
}

func (c ForeignCenter) JinGong() {
	fmt.Printf("ForeignCenter %v 进攻\n", c.Name())
}
