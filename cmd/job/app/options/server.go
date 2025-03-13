package options

import (
	"github.com/chhz0/gojob/internal/job"
	genericopts "github.com/chhz0/gojob/pkg/options"
	"github.com/chhz0/gokit/pkg/cli"
	"github.com/spf13/pflag"
)

type ServerOptions struct {
	Server  string                    `json:"server" mapstructure:"server"`
	OpenTLS bool                      `json:"open_tls" mapstructure:"open_tls"`
	MySQL   *genericopts.MySQLOptions `json:"mysql" mapstructure:"mysql"`
	Redis   *genericopts.RedisOptions `json:"redis" mapstructure:"redis"`
	Job     *JobOptions               `json:"job" mapstructure:"job"`
}

// LocalFlags implements cli.Flager.
func (s *ServerOptions) LocalFlags(fs *pflag.FlagSet) *cli.FlagSet {
	fs.StringVar(&s.Server, "server", "127.0.0.1:4568", "server addr")
	fs.BoolVar(&s.OpenTLS, "open-tls", false, "open tls")

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

	return nil
}

func (s *ServerOptions) Config() (*job.Config, error) {
	cfg := &job.Config{
		MySQL: s.MySQL,
	}

	return cfg, nil
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		Server:  "127.0.0.1:4568",
		OpenTLS: false,
		MySQL:   genericopts.NewMySQLOptions(),
		Redis:   genericopts.NewRedisOptions(),
		Job:     newJobOptions(),
	}
}
