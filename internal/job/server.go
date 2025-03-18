package job

import (
	"github.com/chhz0/gojob/internal/job/router"
	genericopts "github.com/chhz0/gojob/pkg/options"
	"github.com/chhz0/gokit/pkg/log"
	"github.com/chhz0/gokit/pkg/server"
	"github.com/chhz0/gokit/pkg/server/engines"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Mode    string
	Engine  string
	Addr    string
	OpenTLS bool
	MySQL   *genericopts.MySQLOptions
	Redis   *genericopts.RedisOptions
}

type Server struct {
	cfg *Config
	srv server.Server
}

func (cfg *Config) NewServer() (*Server, error) {
	g := useGin()
	srv := server.NewHttp(
		&server.HttpConfig{
			Addr: cfg.Addr,
		},
		engines.Gin(g),
	)

	return &Server{
		cfg: cfg,
		srv: srv,
	}, nil
}

func useGin() *gin.Engine {
	g := gin.New()
	router.Register(g)

	return g
}

func (s *Server) Run() error {
	log.Info("Start to listening the incoming requests at http address", log.String("address", s.cfg.Addr))
	if err := s.srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
