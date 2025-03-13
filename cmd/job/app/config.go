package app

import "github.com/chhz0/gokit/pkg/config"

func vconfig() *config.VConfig {
	return config.NewWith(
		config.WithConfig(&config.LocalConfig{
			ConfigName:  "job-server",
			ConfigType:  "yaml",
			ConfigPaths: []string{
				"../../configs",
			},
		}),
	)
}
