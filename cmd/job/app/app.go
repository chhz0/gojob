package app

import (
	"context"
	"fmt"

	"github.com/chhz0/gojob/cmd/job/app/options"
	"github.com/chhz0/gokit"
	"github.com/chhz0/gokit/pkg/cli"
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
			fmt.Println("go job is running now...")
			fmt.Printf("svr opts: %v\n", svrOpts)
			return nil
		},
		Flager: svrOpts,
	},
		cli.EnableConfig(vc.V()),
		cli.SetConfigHandler(vc.Load),
	)
}
