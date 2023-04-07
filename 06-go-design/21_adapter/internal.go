package adapter

import "fmt"

// 中锋
type Center struct {
	name string
}

func NewCenter(name string) *Center {
	return &Center{
		name: name,
	}
}

// 国内中锋
func (c Center) Name() string {
	return c.name
}

func (c Center) Attack() {
	fmt.Printf("Center %v attack\n", c.Name())
}
