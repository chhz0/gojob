package options

import (
	"time"

	"github.com/chhz0/gokit/pkg/cli"
	"github.com/spf13/pflag"
)

type MySQLOptions struct {
	Dsn             string        `json:"dsn" mapstructure:"dsn"`
	Addr            string        `json:"addr" mapstructure:"addr"`
	Port            string        `json:"port" mapstructure:"port"`
	User            string        `json:"user" mapstructure:"user"`
	Password        string        `json:"password" mapstructure:"password"`
	Database        string        `json:"db" mapstructure:"db"`
	MaxIdleConn     int           `json:"max_idle_conn" mapstructure:"max_idle_conn"`
	MaxActiveConn   int           `json:"max_active_conn" mapstructure:"max_active_conn"`
	MaxConnLifeTime time.Duration `json:"max_conn_life_time" mapstructure:"max_conn_life_time"`
	LogFile         string        `json:"log_file" mapstructure:"log_file"`
}

// LocalFlags implements cli.Flager.
func (m *MySQLOptions) LocalFlags(fs *pflag.FlagSet) *cli.FlagSet {
	fs.StringVar(&m.Dsn, "mysql.dsn", m.Dsn, "mysql dsn")
	fs.StringVar(&m.Addr, "mysql.addr", m.Addr, "mysql addr")
	fs.StringVar(&m.Port, "mysql.port", m.Port, "mysql port")
	fs.StringVar(&m.User, "mysql.user", m.User, "mysql user")
	fs.StringVar(&m.Password, "mysql.password", m.Password, "mysql password")
	fs.StringVar(&m.Database, "mysql.db", m.Database, "mysql database")
	fs.IntVar(&m.MaxIdleConn, "mysql.max-idle-conn", m.MaxIdleConn, "mysql max idle conn")
	fs.IntVar(&m.MaxActiveConn, "mysql.max-active-conn", m.MaxActiveConn, "mysql max active conn")
	fs.DurationVar(&m.MaxConnLifeTime, "mysql.max-idle-time", m.MaxConnLifeTime, "mysql max idle time")
	fs.StringVar(&m.LogFile, "mysql.log-file", m.LogFile, "mysql log file")

	return &cli.FlagSet{
		PFlags:   fs,
		Required: []string{},
	}
}

// PersistentFlags implements cli.Flager.
func (m *MySQLOptions) PersistentFlags(fs *pflag.FlagSet) *cli.FlagSet {
	return &cli.FlagSet{
		PFlags:   fs,
		Required: []string{},
	}
}

var _ cli.Flager = (*MySQLOptions)(nil)

func (m *MySQLOptions) Validate() error {
	return nil
}

func newMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Dsn:             "",
		Addr:            "127.0.0.1",
		Port:            "3306",
		User:            "root",
		Password:        "",
		Database:        "",
		MaxIdleConn:     100,
		MaxActiveConn:   100,
		MaxConnLifeTime: 10 * time.Second,
		LogFile:         "",
	}
}
