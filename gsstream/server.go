package main

type Config struct {
	ListenAddr string
}

type Server struct {
	*Config
	Store Storer
}

func NewServer(cfg *Config) (*Server, error) {
	return &Server{Config: cfg}, nil
}

func (s *Server) Fetch(offset int) ([]byte, error) {
	return s.Store.Fetch(offset)
}
