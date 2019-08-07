package simplefactory

type Producer interface {
	Do()
}

func NewProduct(key string) Producer {
	switch key {
	case "p1":
		return &P1{}
	case "p2":
		return &P2{}
	}
	return nil
}

type P1 struct {
}

func (*P1) Do() {

}

type P2 struct {
}

func (*P2) Do() {

}
