package job

import (
	"fmt"

	genericopts "github.com/chhz0/gojob/pkg/options"
)

type Config struct {
	MySQL *genericopts.MySQLOptions
}

type Server struct {
	cfg *Config
}

func (cfg *Config) NewServer() (*Server, error) {
	return &Server{
		cfg: cfg,
	}, nil
}

func (s *Server) Run() error {
	fmt.Printf("go job is running now...\n")
	fmt.Printf("svr opts: %v\n", s.cfg.MySQL.Addr)

	return nil
}
