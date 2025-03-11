package job

import (
	"context"

	fgcli "github.com/chhz0/framego/base/fg-cli"
)

func NewJober() (*fgcli.Cli, error) {

	return fgcli.NewCli(&fgcli.SimpleCommand{
		CmdName: "job",
		CmdLong: "Job starts an http service and provides an API interface to manage tasks",
		RunFunc: func(ctx context.Context, args []string) error {
			return run(ctx)
		},
	})
}

func run(ctx context.Context) error {
	return nil
}
