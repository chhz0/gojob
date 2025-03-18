package main

import (
	"context"

	"github.com/chhz0/gotasks/cmd/tasks/app"

	_ "go.uber.org/automaxprocs"
)

func main() {
	cmd, err := app.NewtasksCommand()
	if err != nil {
		panic(err)
	}

	if err := cmd.Execute(context.Background()); err != nil {
		panic(err)
	}
}
