package factory_method

import "testing"

func TestFactory(t *testing.T) {
	var factory Factorier

	// 工厂 P1Factory 负责生产 P1
	factory = &P1Factory{}
	p1 := factory.Create()
	if p1.Do() != "P1" {
		t.Fatal("expect P1 but not")
	}

	// 工厂 P2Factory 负责生产 P2
	factory = &P2Factory{}
	p2 := factory.Create()
	if p2.Do() != "P2" {
		t.Fatal("expect P2 but not")
	}

	/**
		通过示例,可以看出,通过这种方式,在调用的时候,只需要知道自己需要什么就可以,不需要像简单工厂模式一样传入一个参数,这样的好处是避免了传入非法参数引起的错误.
	 */
}
