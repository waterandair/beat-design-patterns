package singleton_pattern

func NewSingletonThreadUnsafe(new func() interface{}) *singletonThreadUnsafe {
	return &singletonThreadUnsafe{
		new,
	}
}

var instanceUnsafe interface{}

type singletonThreadUnsafe struct {
	new func() interface{}
}

func (s *singletonThreadUnsafe) GetInstance() interface{} {
	if instanceUnsafe == nil {
		instanceUnsafe = s.new()
	}

	return instanceUnsafe
}
