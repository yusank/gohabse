package gohbase

type Option func(*option)

type option struct {
	addr string
}

func newOption() *option {
	return &option{
		addr: "",
	}
}

func (o *option) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func Addr(addr string) Option {
	return func(o *option) {
		o.addr = addr
	}
}
