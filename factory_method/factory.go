package factory_method

// 工厂接口
type Factorier interface {
	Create() Producer
}

// 产品接口
type Producer interface {
	Do() string
}
