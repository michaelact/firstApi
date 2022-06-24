package log

import (
	"os"

	"github.com/rs/zerolog"

	"github.com/michaelact/firstApi/model/environment"
	"github.com/michaelact/firstApi/helper"
)

func NewLogger(c *environment.Global) *zerolog.Logger {
	var logger zerolog.Logger
	if c.Log.Output == "file" {
		f, err := os.OpenFile(c.Log.OutputFile, os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
		helper.PanicIfError(err)

		logger = zerolog.New(f)
	} else {
		logger = zerolog.New(os.Stdout)
	}

	logger.Level(c.Log.Level)
	logger = logger.With().Timestamp().Logger()
	return &logger
}
