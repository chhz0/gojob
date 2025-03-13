package main

import (
	"context"

	"github.com/chhz0/gojob/cmd/job/app"

	_ "go.uber.org/automaxprocs"
)

func main() {
	cmd, err := app.NewJobCommand()
	if err != nil {
		panic(err)
	}

	if err := cmd.Execute(context.Background()); err != nil {
		panic(err)
	}
}
