package log

import (
	"os"

	"github.com/rs/zerolog"

	"github.com/michaelact/firstApi/model/config"
)

func NewLogger(c *config.Global) *zerolog.Logger { 
	var logger zerolog.Logger
	logger = zerolog.New(os.Stdout)
	logger.Level(c.Log.Level)
	logger = logger.With().Timestamp().Logger() 

	return &logger
}
