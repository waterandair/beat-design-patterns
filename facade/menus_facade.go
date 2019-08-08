package facade

// 套餐方式
type MenuFacader interface {
	OrderPackage(name string) []string
}

type MenuFacade struct {
	Menuer
}

var _ MenuFacader = (*MenuFacade)(nil)

func NewMenuFacade() MenuFacader {
	return &MenuFacade{}
}

func (m *MenuFacade) OrderPackage(name string) []string {
	orders := make([]string, 0)
	switch name {
	case "A":
		return append(orders, m.OrderBeef(), m.OrderColdDish(), m.OrderDumplings(), m.OrderHotPot())
	case "B":
		return append(orders, m.OrderBeef(), m.OrderColdDish(), m.OrderDumplings(), m.OrderHotPot(), m.OrderBear())
	default:
		return orders
	}
}
