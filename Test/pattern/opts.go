package pattern

// 选项模式
type Opts struct {
	maxConn int
	id      string
	tls     bool
}

type OptFunc func() *Opts

type Server struct {
	Opts
}

func WithMaxConn(maxConn int) OptFunc {
	return func() *Opts {
		return &Opts{
			maxConn: maxConn,
		}
	}
}

func WithID(id string) OptFunc {
	return func() *Opts {
		return &Opts{
			id: id,
		}
	}
}

func WithTls(tls bool) OptFunc {
	return func() *Opts {
		return &Opts{
			tls: tls,
		}
	}
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "1111",
		tls:     false,
	}
}

func NewServer(optfuncs ...OptFunc) *Server {
	o := defaultOpts()
	for _, optfunc := range optfuncs {
		optfunc()
	}
	return &Server{
		Opts: o,
	}
}
