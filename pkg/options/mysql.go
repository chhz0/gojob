package genericopts

import (
	"fmt"
	"time"

	"github.com/chhz0/gokit/pkg/cli"
	"github.com/spf13/pflag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLOptions struct {
	Addr            string        `json:"addr" mapstructure:"addr"`
	Port            string        `json:"port" mapstructure:"port"`
	User            string        `json:"user" mapstructure:"user"`
	Password        string        `json:"password" mapstructure:"password"`
	Database        string        `json:"db" mapstructure:"db"`
	MaxIdleConn     int           `json:"max_idle_conn" mapstructure:"max_idle_conn"`
	MaxOpenConn     int           `json:"max_open_conn" mapstructure:"max_open_conn"`
	MaxConnLifeTime time.Duration `json:"max_conn_life_time" mapstructure:"max_conn_life_time"`
	LogFile         string        `json:"log_file" mapstructure:"log_file"`
}

// LocalFlags implements cli.Flager.
func (m *MySQLOptions) LocalFlags(fs *pflag.FlagSet) *cli.FlagSet {
	fs.StringVar(&m.Addr, "mysql.addr", m.Addr, "mysql addr")
	fs.StringVar(&m.Port, "mysql.port", m.Port, "mysql port")
	fs.StringVar(&m.User, "mysql.user", m.User, "mysql user")
	fs.StringVar(&m.Password, "mysql.password", m.Password, "mysql password")
	fs.StringVar(&m.Database, "mysql.db", m.Database, "mysql database")
	fs.IntVar(&m.MaxIdleConn, "mysql.max-idle-conn", m.MaxIdleConn, "mysql max idle conn")
	fs.IntVar(&m.MaxOpenConn, "mysql.max-open-conn", m.MaxOpenConn, "mysql max open conn")
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

func (m *MySQLOptions) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.User,
		m.Password,
		m.Addr,
		m.Port,
		m.Database,
	)
}

func (m *MySQLOptions) NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(m.MaxOpenConn)
	sqlDB.SetMaxIdleConns(m.MaxIdleConn)
	sqlDB.SetConnMaxIdleTime(m.MaxConnLifeTime)

	return db, nil
}

func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Addr:            "127.0.0.1",
		Port:            "3306",
		User:            "root",
		Password:        "",
		Database:        "",
		MaxIdleConn:     100,
		MaxOpenConn:     100,
		MaxConnLifeTime: 10 * time.Second,
		LogFile:         "",
	}
}
