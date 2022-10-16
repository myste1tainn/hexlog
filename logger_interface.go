package log

type LoggerInterface interface {
	DebugJsonf(json any, format string, args ...interface{})
	InfofJson(json any, format string, args ...interface{})
	PrintJsonf(json any, format string, args ...interface{})
	WarnfJson(json any, format string, args ...interface{})
	ErrorJsonf(json any, format string, args ...interface{})
	FatalJsonf(json any, format string, args ...interface{})
	PanicJsonf(json any, format string, args ...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})

	AddPayload(func() map[string]any)
}
