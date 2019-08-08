package facade

// 菜单方式
type Menuer interface {
	OrderBeef() string
	OrderDumplings() string
	OrderHotPot() string
	OrderColdDish() string
	OrderBear() string
}

var _ Menuer = (*Menu)(nil)

func NewMenu() Menuer {
	return &Menu{}
}

type Menu struct {
}

func (m *Menu) OrderBeef() string {
	return "Beef"
}

func (m *Menu) OrderDumplings() string {
	return "Dumplings"
}

func (m *Menu) OrderHotPot() string {
	return "HotPot"
}

func (m *Menu) OrderColdDish() string {
	return "ColdDish"
}

func (m *Menu) OrderBear() string {
	return "Bear"
}
