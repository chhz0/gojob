package app

import (
	"context"

	"github.com/chhz0/gojob/cmd/job/app/options"
	"github.com/chhz0/gojob/internal/pkg/logger"
	"github.com/chhz0/gokit"
	"github.com/chhz0/gokit/pkg/cli"
	"github.com/chhz0/gokit/pkg/log"
)

func NewJobCommand() (cli.CliExector, error) {
	vc := vconfig()
	svrOpts := options.NewServerOptions()
	return gokit.NewCli(&cli.SimpleCommand{
		CmdName:  "job",
		CmdShort: "job is a simple asynchronous task processing framework",
		CmdLong: `job is an asynchronous task processing framework driven by the Go language,
supporting task management, custom tasks, etc.`,
		RunFunc: func(ctx context.Context, args []string) error {
			return run(svrOpts)
		},
		Flager: svrOpts,
		Commanders: []cli.Commander{
			newVersion(),
		},
	},
		cli.EnableConfig(vc.V()),
		cli.SetConfigHandler(vc.Load),
	)
}

func run(opts *options.ServerOptions) error {
	logger.NewLogger()
	log.Info("job is starting with the options", log.Any("options", opts))
	if err := opts.Validate(); err != nil {
		return err
	}

	cfg, err := opts.Config()
	if err != nil {
		return err
	}

	svr, err := cfg.NewServer()
	if err != nil {
		return err
	}

	return svr.Run()
}
