package factory_method

type P1Factory struct {
}

func (*P1Factory) Create() Producer {
	return &P1{}
}

type P1 struct {
}

func (*P1) Do() string {
	return "P1"
}
