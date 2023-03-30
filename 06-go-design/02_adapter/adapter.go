package adapter

// NewAdapter 是Adapter的工厂函数, 将Adptee转化为Target
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

// Adapter 是转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee
}

// Request 实现Target接口
func (a *adapter) Request() string {
	return a.SpecificRequest()
}
