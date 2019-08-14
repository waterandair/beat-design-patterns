package factory_method

type P2Factory struct {
}

func (*P2Factory) Create() Producer {
	return &P2{}
}

type P2 struct {
}

func (*P2) Do() string {
	return "P2"
}
