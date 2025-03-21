package app

import (
	"context"

	"github.com/chhz0/gotasks/pkg/version"
	"github.com/chhz0/gokit/pkg/cli"
)

func newVersion() cli.Commander {
	vf := &version.VersionFlags{
		Format: "string",
	}
	return &cli.SimpleCommand{
		CmdName:  "version",
		CmdShort: "Print the version number of tasks",
		CmdLong: `version command prints out the version of tasks.
if use the flag --raw, it will print out all the version information.
{version, gitCommit, gitTreeState, buildDate, goVersion, compiler, platform}`,
		RunFunc: func(ctx context.Context, args []string) error {
			version.PrintVersion(vf.Format)
			return nil
		},
		Flager: vf,
	}
}
