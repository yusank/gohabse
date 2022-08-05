package gohbase

import "github.com/tsuna/gohbase"

type Option func(*option)

type option struct {
	addr        string
	gohbaseOpts []gohbase.Option
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

// GoHbaseOpts is uses github.com/tsuna/gohbase options.
func GoHbaseOpts(opts ...gohbase.Option) Option {
	return func(o *option) {
		if o.gohbaseOpts == nil {
			o.gohbaseOpts = make([]gohbase.Option, 0)
		}
		o.gohbaseOpts = append(o.gohbaseOpts, opts...)
	}
}
