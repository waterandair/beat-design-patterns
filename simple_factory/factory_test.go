package simplefactory

import "testing"

func TestNewProduct(t *testing.T) {
	p1 := NewProduct("p1")
	if _, ok := p1.(*P1); !ok {
		t.Error("expect P1 but not")
	}

	p2 := NewProduct("p2")
	if _, ok := p2.(*P2); !ok {
		t.Error("expect P2 but not")
	}

	p3 := NewProduct("p3")
	if _, ok := p3.(interface{}); !ok {
		t.Error("expect nil interface but not")
	}

}
