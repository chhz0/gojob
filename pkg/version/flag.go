package version

import (
	"fmt"

	"github.com/chhz0/gokit/pkg/cli"
	"github.com/spf13/pflag"
)

type VersionFlags struct {
	Format string
}

// LocalFlags implements cli.Flager.
func (v *VersionFlags) LocalFlags(fs *pflag.FlagSet) *cli.FlagSet {
	return &cli.FlagSet{
		PFlags:   fs,
		Required: []string{},
	}
}

// PersistentFlags implements cli.Flager.
func (v *VersionFlags) PersistentFlags(fs *pflag.FlagSet) *cli.FlagSet {
	fs.StringVar(&v.Format, "raw", v.Format, "print version info in <string | raw | json> format")
	return &cli.FlagSet{
		PFlags:   fs,
		Required: []string{},
	}
}

var _ cli.Flager = (*VersionFlags)(nil)

// todo: 优化version的设置
var versionVal = ""

func Flag(fs *pflag.FlagSet) {
	fs.StringVarP(&versionVal, "version", "v", versionVal, "print version info")
}

func PrintVersion(format string) {
	switch format {
	case "raw":
		fmt.Printf("%s\n", Get().Text())
	case "json":
		fmt.Printf("%s\n", Get().ToJSON())
	default:
		fmt.Printf("%s\n", Get().Text())
	}
}
