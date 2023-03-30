package adapter

// Adaptee 是被适配的目标接口
type Adaptee interface {
	SpecificRequest() string
}

// NewAdaptee 是被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// AdapteeImpl 是被适配的目标类
type adapteeImpl struct{}

// SpecificRequest 是目标类的一个方法
func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}
