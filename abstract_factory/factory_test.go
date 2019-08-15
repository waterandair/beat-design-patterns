package abstract_factory

import (
	"log"
	"testing"
)

var f Themer

func TestSimpleFactory(t *testing.T) {
	f = &SimpleTheme{}
	m := f.CreateMenu()
	w := f.CreateWorkSpace()

	log.Println(m.ShowTabs())
	log.Println(w.ShowTools())
}

func TestClassicFactory(t *testing.T) {
	f = &ClassicTheme{}
	m := f.CreateMenu()
	w := f.CreateWorkSpace()

	log.Println(m.ShowTabs())
	log.Println(w.ShowTools())
}
