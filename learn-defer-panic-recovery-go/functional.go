// an example of functional options patterns in go
package main

type Server struct {
	Address string
	Port    int
}

type Option func(*Server)

func WithAddress(address string) Option {
	return func(s *Server) {
		s.Address = address
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func NewServer(opts ...Option) *Server {
	server := &Server{
		Address: "localhost",
		Port:    8080,
	}

	for _, opt := range opts {
		opt(server)
	}
	return server
}
