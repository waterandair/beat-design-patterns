package abstract_factory

// 在写代码前,要先明白一个关系: 抽象工厂封装工厂接口,具体工厂实现抽象工厂接口,具体工厂生产抽象产品,抽象产品封装产品接口,具体产品实现抽象产品接口.

// 这里以模拟开发一个编辑器举例,现代的编辑器通常都会提供多种主题模式供用户选择,比如简单模式, 经典模式,以后还可能需要提供炫酷模式等等...

// 抽象工厂: 封装工厂接口
type Themer interface {
	CreateMenu() Menuer
	CreateWorkSpace() WorkSpacer
}

// 具体工厂: 实现抽象工厂接口,生产抽象产品 Menuer 和 WorkSpacer
type SimpleTheme struct {
}

func (s *SimpleTheme) CreateMenu() Menuer {
	return &SimpleMenu{}
}

func (s *SimpleTheme) CreateWorkSpace() WorkSpacer {
	return &SimpleWorkSpacer{}
}

// 具体工厂: 生产抽象产品
type ClassicTheme struct {
}

func (c *ClassicTheme) CreateMenu() Menuer {
	return &ClassicMenu{}
}

func (c *ClassicTheme) CreateWorkSpace() WorkSpacer {
	return &ClassWorkSpacer{}
}

// 抽象产品: 封装产品接口
type Menuer interface {
	ShowTabs() []string
}

// 抽象产品: 封装产品接口
type WorkSpacer interface {
	ShowTools() []string
}

// 具体产品: 实现抽象产品接口 Menuer
type SimpleMenu struct {
}

func (*SimpleMenu) ShowTabs() []string {
	return []string{"file"}
}

// 具体产品: 实现抽象产品接口 Menuer
type ClassicMenu struct {
}

func (*ClassicMenu) ShowTabs() []string {
	return []string{"file", "edit", "view", "help"}
}

// 具体产品: 实现抽象产品接口  WorkSpacer
type SimpleWorkSpacer struct {
}

func (*SimpleWorkSpacer) ShowTools() []string {
	return []string{"editor"}
}

// 具体产品: 实现抽象产品接口  WorkSpacer
type ClassWorkSpacer struct {
}

func (*ClassWorkSpacer) ShowTools() []string {
	return []string{"editor", "project", "terminal"}
}
