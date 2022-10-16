package log

import (
	zlog "github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var L *Logger = &Logger{Logger: &zlog.Logger}

func InitGlobalLogger() *Logger {
	var cfg Config
	viper.UnmarshalKey("logger.global", &cfg)
	L = &Logger{
		cfg:            cfg,
		Logger:         newZeroLog(cfg),
		DefaultPayload: map[string]any{},
	}
	return L
}

func NewLogger(cfg Config) *Logger {
	return &Logger{
		cfg:            cfg,
		Logger:         newZeroLog(cfg),
		DefaultPayload: map[string]any{},
	}
}
