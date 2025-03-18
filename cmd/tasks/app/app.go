package app

import (
	"context"

	"github.com/chhz0/gotasks/cmd/tasks/app/options"
	"github.com/chhz0/gotasks/internal/pkg/logger"
	"github.com/chhz0/gokit"
	"github.com/chhz0/gokit/pkg/cli"
	"github.com/chhz0/gokit/pkg/log"
)

func NewtasksCommand() (cli.CliExector, error) {
	vc := vconfig()
	svrOpts := options.NewServerOptions()
	return gokit.NewCli(&cli.SimpleCommand{
		CmdName:  "tasks",
		CmdShort: "tasks is a simple asynchronous task processing framework",
		CmdLong: `tasks is an asynchronous task processing framework driven by the Go language,
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
	log.Info("tasks is starting with the options", log.Any("options", opts))
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
