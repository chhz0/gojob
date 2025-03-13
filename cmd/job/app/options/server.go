package options

import (
	"github.com/chhz0/gokit/pkg/cli"
	"github.com/spf13/pflag"
)

type ServerOptions struct {
	Server  string        `json:"server" mapstructure:"server"`
	OpenTLS bool          `json:"open_tls" mapstructure:"open_tls"`
	MySQL   *MySQLOptions `json:"mysql" mapstructure:"mysql"`
	Redis   *RedisOptions `json:"redis" mapstructure:"redis"`
	Job     *JobOptions   `json:"job" mapstructure:"job"`
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

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		Server:  "127.0.0.1:4568",
		OpenTLS: false,
		MySQL:   newMySQLOptions(),
		Redis:   newRedisOptions(),
		Job:     newJobOptions(),
	}
}
