package facade

import (
	"fmt"
	"testing"
)

// 用菜单的方式点餐
func MenuOrder(t *testing.T) {
	menu := NewMenu()

	orders := make([]string, 0)
	orders = append(orders, menu.OrderBear())
	orders = append(orders, menu.OrderDumplings())
	orders = append(orders, menu.OrderColdDish())
	orders = append(orders, menu.OrderBear())
	fmt.Println(orders)
}

// 用套餐的方式点餐
func PackageOrder(t *testing.T) {
	menu := NewMenuFacade()
	orders := menu.OrderPackage("A")
	fmt.Println(orders)
}

/**
	对比这两种方式可以看出,使用外观模式的方式,不仅可以使调用变得简单,也可以避免调用中发生的误操作.
 */
