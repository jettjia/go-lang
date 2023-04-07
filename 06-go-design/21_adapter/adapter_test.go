package adapter

import "testing"

func Test_Adapter(t *testing.T) {
	c := NewCenter("Internal")
	c.Attack()

	fc := NewForeignCenter("Foreign")
	afc := NewAdapter(fc)
	afc.Attack()
}
