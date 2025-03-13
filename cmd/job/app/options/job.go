package options

import (
	"time"

	"github.com/chhz0/gokit/pkg/cli"
	"github.com/spf13/pflag"
)

type JobOptions struct {
	MaxTableRows        int           `json:"max_table_rows" mapstructure:"max_table_rows"`
	SplitInterval       time.Duration `json:"split_interval" mapstructure:"split_interval"`
	LongProcessInterval time.Duration `json:"long_process_interval" mapstructure:"long_process_interval"`
	CheckInterval       time.Duration `json:"check_interval" mapstructure:"check_interval"`
	MaxProcessTime      time.Duration `json:"max_process_time" mapstructure:"max_process_time"`
	OpenGovern          bool          `json:"open_govern" mapstructure:"open_govern"`
}

// LocalFlags implements cli.Flager.
func (j *JobOptions) LocalFlags(fs *pflag.FlagSet) *cli.FlagSet {
	fs.IntVar(&j.MaxTableRows, "job.max-table-rows", j.MaxTableRows, "max table rows")
	fs.DurationVar(&j.SplitInterval, "job.split-interval", j.SplitInterval, "split interval")
	fs.DurationVar(&j.LongProcessInterval, "job.long-process-interval", j.LongProcessInterval, "long process interval")
	fs.DurationVar(&j.CheckInterval, "job.check-interval", j.CheckInterval, "check interval")
	fs.DurationVar(&j.MaxProcessTime, "job.max-process-time", j.MaxProcessTime, "max process time")
	fs.BoolVar(&j.OpenGovern, "job.open-govern", j.OpenGovern, "open govern")

	return &cli.FlagSet{
		PFlags:   fs,
		Required: []string{},
	}
}

// PersistentFlags implements cli.Flager.
func (j *JobOptions) PersistentFlags(fs *pflag.FlagSet) *cli.FlagSet {
	return &cli.FlagSet{
		PFlags:   fs,
		Required: []string{},
	}
}

var _ cli.Flager = (*JobOptions)(nil)

func newJobOptions() *JobOptions {
	return &JobOptions{
		MaxTableRows:        5000000,
		SplitInterval:       30 * time.Second,
		LongProcessInterval: 10 * time.Second,
		CheckInterval:       10 * time.Second,
		MaxProcessTime:      10 * time.Second,
		OpenGovern:          false,
	}
}
