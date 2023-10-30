package main

import "fmt"

type Opts struct {
	maxConn int
	id string
	tls bool
}

type OptFunc func(*Opts)

func defaultOpts() *Opts {
	return &Opts{
		maxConn: 333,
		id: "defualt",
		tls: false,
	}
}

func withTls(opts *Opts) {
	opts.tls = true
}

func withID(s string) OptFunc {
	return func(opts *Opts) {
		opts.id = s
	}
}

type Server struct {
	Opts
}

func withMaxConn(n int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = n
	}
}

func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(o)
	}
	return &Server{
		Opts: *o,
	}
}

func main() {
	s := newServer(withTls, withMaxConn(99), withID("fdasdfasd"))
	fmt.Printf("%v", s)
}



