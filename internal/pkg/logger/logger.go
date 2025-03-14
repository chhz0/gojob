package logger

import (
	"io"
	"os"

	"github.com/chhz0/gokit/pkg/log"
)

func NewLogger() {
	opts := []log.ZapOption{
		log.WithCaller(true),
		log.AddCallerSkip(3),
		// log.AddCaller(),
	}

	l := log.NewLogger(
		logOutput,
		log.InfoLevel,
		log.JsonEncoder,
		opts...,
	)

	log.ReplaceDefault(l)
}

func logOutput() io.Writer {
	return os.Stdout
}
