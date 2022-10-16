package log

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type AddPayloadHandler = func() map[string]any

type Logger struct {
	*zerolog.Logger

	// cfg is the configuration this Logger was created with
	cfg Config

	DefaultPayload map[string]any
	addPayloads    []AddPayloadHandler

	// parentLogger who instantiated this logger, nil if the logger is Global
	parentLogger *Logger

	// chileLoggers instantiate by this logger
	childLoggers []*Logger
}

func NewParentLogger() *Logger {
	return L.NewChildLogger()
}

func (l *Logger) NewChildLogger() *Logger {
	logger := NewLogger(l.cfg)
	if l.DefaultPayload == nil {
		logger.DefaultPayload = map[string]any{}
	} else {
		p := map[string]any{}
		for k, v := range l.DefaultPayload {
			p[k] = v
		}
		logger.DefaultPayload = p
	}
	logger.addPayloads = l.addPayloads
	logger.parentLogger = l
	l.childLoggers = append(l.childLoggers, logger)
	return logger
}

func (l *Logger) NewSpanLogger() *Logger {
	logger := l.NewChildLogger()
	currSpanId := l.DefaultPayload["spanId"]
	if currSpanId != "" && currSpanId != nil {
		logger.DefaultPayload["spanId"] = fmt.Sprintf("%s/%s", currSpanId, newSpanId())
	} else {
		logger.DefaultPayload["spanId"] = newSpanId()
	}

	return logger
}

func newSpanId() string {
	uuidString := uuid.New().String()
	components := strings.Split(uuidString, "-")
	if len(components) > 1 {
		// this is to use the short version of the string
		return components[0]
	} else {
		// otherwise just use the long one
		return uuidString
	}
}

func (l *Logger) GetChildLoggerCount() int {
	return len(l.childLoggers)
}

func (l *Logger) getArgs(args ...interface{}) []interface{} {
	if len(l.DefaultPayload) > 0 {
		args = append(args, l.DefaultPayload)
	}
	for _, fn := range l.addPayloads {
		m := fn()
		if len(m) > 0 {
			args = append(args, m)
		}
	}
	return args
}

func (l *Logger) AddPayload(addPayload AddPayloadHandler) {
	l.addPayloads = append(l.addPayloads, addPayload)
}

func (l *Logger) Destroy() {
	if l.parentLogger == nil {
		return
	}
	if l.parentLogger.childLoggers == nil {
		return
	}

	childLoggers := l.parentLogger.childLoggers

	index := -1
	for i, logger := range childLoggers {
		if l == logger {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	childLoggers[index] = nil
	if index == len(l.parentLogger.childLoggers)-1 {
		childLoggers = childLoggers[:index]
	} else if index == 0 {
		childLoggers = childLoggers[1:]
	} else {
		childLoggers = append(childLoggers[:index], childLoggers[index+1:]...)
	}
	l.parentLogger.childLoggers = childLoggers
}
