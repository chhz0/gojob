package options

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	"github.com/chhz0/gojob/internal/job"
	genericopts "github.com/chhz0/gojob/pkg/options"
	"github.com/chhz0/gokit/pkg/cli"
	"github.com/spf13/pflag"
)

type ServerOptions struct {
	Mode    string `json:"mode" mapstructure:"mode"`
	Engine  string `json:"engine" mapstructure:"engine"`
	Addr    string `json:"addr" mapstructure:"addr"`
	OpenTLS bool   `json:"open_tls" mapstructure:"open_tls"`

	MySQL *genericopts.MySQLOptions `json:"mysql" mapstructure:"mysql"`
	Redis *genericopts.RedisOptions `json:"redis" mapstructure:"redis"`
	Job   *JobOptions               `json:"job" mapstructure:"job"`
}

// LocalFlags implements cli.Flager.
func (s *ServerOptions) LocalFlags(fs *pflag.FlagSet) *cli.FlagSet {
	fs.StringVar(&s.Mode, "mode", s.Mode, "server mode (dev | release)")
	fs.StringVar(&s.Engine, "engine", s.Engine, "engine (gin | grpc | ...http)")
	fs.StringVar(&s.Addr, "addr", s.Addr, "server addr")
	fs.BoolVar(&s.OpenTLS, "open-tls", s.OpenTLS, "open tls")

	s.MySQL.LocalFlags(fs)
	s.Redis.LocalFlags(fs)

	return &cli.FlagSet{
		PFlags:   fs,
		Required: []string{},
	}
}

// PersistentFlags implements cli.Flager.
func (s *ServerOptions) PersistentFlags(fs *pflag.FlagSet) *cli.FlagSet {
	return &cli.FlagSet{
		PFlags:   fs,
		Required: []string{},
	}
}

var _ cli.Flager = (*ServerOptions)(nil)

func (s *ServerOptions) Validate() error {
	if err := s.MySQL.Validate(); err != nil {
		return err
	}

	if err := s.Redis.Validate(); err != nil {
		return err
	}

	if s.Addr == "" {
		return fmt.Errorf("server address can not be empty")
	}
	_, portStr, err := net.SplitHostPort(s.Addr)
	if err != nil {
		return fmt.Errorf("server address is invalid")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("server port is invalid")
	}

	return nil
}

func (s *ServerOptions) ToJSON() string {
	j, _ := json.Marshal(s)
	return string(j)
}

func (s *ServerOptions) Config() (*job.Config, error) {
	cfg := &job.Config{
		Mode:    s.Mode,
		Engine:  s.Engine,
		Addr:    s.Addr,
		OpenTLS: s.OpenTLS,
		MySQL:   s.MySQL,
		Redis:   s.Redis,
	}

	return cfg, nil
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		Mode:    "dev",
		Engine:  "gin",
		Addr:    "127.0.0.1:4568",
		OpenTLS: false,
		MySQL:   genericopts.NewMySQLOptions(),
		Redis:   genericopts.NewRedisOptions(),
		Job:     newJobOptions(),
	}
}
