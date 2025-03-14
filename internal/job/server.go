package job

import (
	"net/http"
	"time"

	genericopts "github.com/chhz0/gojob/pkg/options"
	"github.com/chhz0/gokit/pkg/log"
	"github.com/chhz0/gokit/pkg/server"
	"github.com/chhz0/gokit/pkg/server/engines"
	"github.com/gin-gonic/gin"
)

type Config struct {
	ServerMode string
	Addr       string
	OpenTLS    bool
	MySQL      *genericopts.MySQLOptions
}

type Server struct {
	cfg *Config
	srv server.Server
}

func (cfg *Config) NewServer() (*Server, error) {

	g := useGin()

	srv := server.NewHttp(
		&server.HttpConfig{
			Addr:         cfg.Addr,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
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

	g.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Page not found.",
		})
	})

	g.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "ok",
		})
	})

	return g
}

func (s *Server) Run() error {
	log.Info("Start to listening the incoming requests at http address", log.String("address", s.cfg.Addr))
	if err := s.srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
