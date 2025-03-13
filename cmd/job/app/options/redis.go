package options

import (
	"github.com/chhz0/gokit/pkg/cli"
	"github.com/spf13/pflag"
)

type RedisOptions struct {
	Url      string `json:"url" mapstructure:"url"`
	User     string `json:"user" mapstructure:"user"`
	Passwrod string `json:"password" mapstructure:"password"`
	DB int    `json:"db" mapstructure:"db"`
}

// LocalFlags implements cli.Flager.
func (r *RedisOptions) LocalFlags(fs *pflag.FlagSet) *cli.FlagSet {
	fs.StringVar(&r.Url, "redis.url", r.Url, "redis url")
	fs.StringVar(&r.User, "redis.user", r.User, "redis user")
	fs.StringVar(&r.Passwrod, "redis.password", r.Passwrod, "redis password")
	fs.IntVar(&r.DB, "redis.db", r.DB, "redis database")

	return &cli.FlagSet{
		PFlags:   fs,
		Required: []string{},
	}
}

// PersistentFlags implements cli.Flager.
func (r *RedisOptions) PersistentFlags(fs *pflag.FlagSet) *cli.FlagSet {
	return &cli.FlagSet{
		PFlags:   fs,
		Required: []string{},
	}
}

var _ cli.Flager = (*RedisOptions)(nil)

func (r *RedisOptions) Validate() error {
	return nil
}

func newRedisOptions() *RedisOptions {
	return &RedisOptions{
		Url:      "127.0.0.1:6379",
		User:     "",
		Passwrod: "",
		DB: 0,
	}
}
