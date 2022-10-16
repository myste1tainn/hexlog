package log

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog"
)

func (l *Logger) printJson(event *zerolog.Event, json any, format string, args ...any) {
	putLogToEvent(event, json)
	l.addDefaults(event)
	msg := fmt.Sprintf(format, args...)
	event.
		Str(GcpJsonLoggingRawMessageKey, format).
		Str(GcpJsonLoggingMessageKey, msg).
		Msg(msg)
}

func (l *Logger) DebugJsonf(json any, format string, args ...interface{}) {
	l.printJson(l.Logger.Debug(), json, format, args...)
}

func (l *Logger) InfoJsonf(json any, format string, args ...interface{}) {
	l.printJson(l.Logger.Info(), json, format, args...)
}

func (l *Logger) PrintJsonf(json any, format string, args ...interface{}) {
	l.printJson(l.Logger.Info(), json, format, args...)
}

func (l *Logger) WarnJsonf(json any, format string, args ...interface{}) {
	l.printJson(l.Logger.Warn(), json, format, args...)
}

func (l *Logger) ErrorJsonf(json any, format string, args ...interface{}) {
	l.printJson(l.Logger.Error(), json, format, args...)
}

func (l *Logger) FatalJsonf(json any, format string, args ...interface{}) {
	l.printJson(l.Logger.Fatal(), json, format, args...)
}

func (l *Logger) PanicJsonf(json any, format string, args ...interface{}) {
	l.printJson(l.Logger.Panic(), json, format, args...)
}

func (l *Logger) addDefaults(event *zerolog.Event) {
	if len(l.DefaultPayload) > 0 {
		for k, v := range l.DefaultPayload {
			event.Str(k, fmt.Sprintf("%v", v))
		}
	}

	if len(l.addPayloads) > 0 {
		for _, fn := range l.addPayloads {
			m := fn()
			if m != nil {
				for k, v := range m {
					event.Str(k, fmt.Sprintf("%v", v))
				}
			}
		}
	}
}

func (l *Logger) printfOut(event *zerolog.Event, format string, args ...any) {
	l.addDefaults(event)
	msg := fmt.Sprintf(format, args...)
	event.
		Str(GcpJsonLoggingRawMessageKey, format).
		Str(GcpJsonLoggingMessageKey, msg).
		Msg(msg)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.printfOut(l.Logger.Debug(), format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.printfOut(l.Logger.Info(), format, args...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.printfOut(l.Logger.Info(), format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.printfOut(l.Logger.Warn(), format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.printfOut(l.Logger.Error(), format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.printfOut(l.Logger.Fatal(), format, args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.printfOut(l.Logger.Panic(), format, args...)
}

func (l *Logger) printOut(event *zerolog.Event, args ...any) {
	l.addDefaults(event)
	strs := []string{}
	for _, v := range args {
		strs = append(strs, fmt.Sprintf("%v", v))
	}
	msg := strings.Join(strs, ",")
	event.
		Str(GcpJsonLoggingMessageKey, msg).
		Msg(msg)
}

func (l *Logger) Debug(args ...interface{}) {
	l.printOut(l.Logger.Debug(), args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.printOut(l.Logger.Info(), args...)
}

func (l *Logger) Print(args ...interface{}) {
	l.printOut(l.Logger.Info(), args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.printOut(l.Logger.Warn(), args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.printOut(l.Logger.Error(), args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.printOut(l.Logger.Fatal(), args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.printOut(l.Logger.Panic(), args...)
}

func (l *Logger) Debugln(args ...interface{}) {
	l.printOut(l.Logger.Debug(), args...)
}

func (l *Logger) Infoln(args ...interface{}) {
	l.printOut(l.Logger.Info(), args...)
}

func (l *Logger) Println(args ...interface{}) {
	l.printOut(l.Logger.Info(), args...)
}

func (l *Logger) Warnln(args ...interface{}) {
	l.printOut(l.Logger.Warn(), args...)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.printOut(l.Logger.Error(), args...)
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.printOut(l.Logger.Fatal(), args...)
}

func (l *Logger) Panicln(args ...interface{}) {
	l.printOut(l.Logger.Panic(), args...)
}

func putLogToEventf(event *zerolog.Event, format string, args ...any) string {
	putLogToEvent(event, args...)
	return fmt.Sprintf(format, args...)
}

func putLogToEvent(event *zerolog.Event, args ...any) string {
	for i, arg := range args {
		key := fmt.Sprintf("%d", i)
		switch t := arg.(type) {
		case map[string]any:
			for k, v := range t {
				putValueToEvent(event, k, v)
			}
		case []any:
			for i, v := range t {
				putValuesToEvent(event, i, v)
			}
		case string, int, bool, int64, int32, float32, float64:
			event.Str(key, fmt.Sprintf("%v", t))
		default:
			fuzzyPutValueToEvent(event, key, t)
		}
	}
	return ""
}

func putValueToEvent(event *zerolog.Event, key string, val any) {
	switch t := val.(type) {
	case map[string]any:
		event.Dict(key, createDict(t))
	case []any:
		event.Array(key, createArrayFromAny(t))
	case []map[string]any:
		event.Array(key, createArrayFromMap(t))
	case string, int, bool, int64, int32, float32, float64:
		event.Str(key, fmt.Sprintf("%v", t))
	default:
		fuzzyPutValueToEvent(event, key, t)
	}
}
func putValuesToEvent(event *zerolog.Event, index int, val any) {
	key := fmt.Sprintf("%d", index)
	switch t := val.(type) {
	case map[string]any:
		event.Dict(key, createDict(t))
	case []any:
		event.Array(key, createArrayFromAny(t))
	case []map[string]any:
		event.Array(key, createArrayFromMap(t))
	case string, int, bool, int64, int32, float32, float64:
		event.Str(key, fmt.Sprintf("%v", t))
	default:
		fuzzyPutValueToEvent(event, key, t)
	}
}
