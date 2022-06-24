package server

type Options struct {
	Path string
}

func NewServerOptions() *Options {
	s := Options{}
	return &s
}
