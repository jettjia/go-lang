package adapter

// IPlayer是球员接口，定义Attack方法，普通中锋实现Attack接口，外籍中锋没有提供Attack接口，于是需要一个适配器来适配

// 球员接口 受改造
type Iplayer interface {
	Name() string
	Attack()
}
