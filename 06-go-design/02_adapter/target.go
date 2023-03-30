package adapter

// Target 是适配的目标接口
type Target interface {
	Request() string
}
