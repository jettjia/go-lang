package adapter

// Adapter 改造国外中锋
type Adapter struct {
	fc *ForeignCenter
}

func NewAdapter(fc *ForeignCenter) *Adapter {
	return &Adapter{
		fc: fc,
	}
}

func (a Adapter) Name() string {
	return a.fc.Name()
}

func (a Adapter) Attack() {
	a.fc.JinGong()
}
