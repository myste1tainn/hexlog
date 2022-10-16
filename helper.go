package log

import (
	"os"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func newZeroLog(cfg Config) *zerolog.Logger {
	// This is the setup so these values are consumable by GCP
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	zerolog.TimestampFieldName = "timestamp"
	zerolog.LevelFieldName = "severity"
	zerolog.MessageFieldName = "message"

	lvl, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(lvl)
	}

	if ParseInfra(cfg.Infra) == InfraLocal {
		zlog.Logger = zlog.With().Timestamp().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stdout})
		return &zlog.Logger
	}

	return &logger
}
