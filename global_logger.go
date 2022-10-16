package log

func DebugJsonf(obj any, format string, args ...interface{}) {
	L.DebugJsonf(obj, format, args...)
}

func InfoJsonf(obj any, format string, args ...interface{}) {
	L.InfoJsonf(obj, format, args...)
}

func PrintJsonf(obj any, format string, args ...interface{}) {
	L.PrintJsonf(obj, format, args...)
}

func WarnJsonf(obj any, format string, args ...interface{}) {
	L.WarnJsonf(obj, format, args...)
}

func ErrorJsonf(obj any, format string, args ...interface{}) {
	L.ErrorJsonf(obj, format, args...)
}

func FatalJsonf(obj any, format string, args ...interface{}) {
	L.FatalJsonf(obj, format, args...)
}

func PanicJsonf(obj any, format string, args ...interface{}) {
	L.PanicJsonf(obj, format, args...)
}

func Debugf(format string, args ...interface{}) {
	L.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	L.Infof(format, args...)
}

func Printf(format string, args ...interface{}) {
	L.Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	L.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	L.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	L.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	L.Panicf(format, args...)
}

func Debug(args ...interface{}) {
	L.Debug(args)
}

func Info(args ...interface{}) {
	L.Info(args)
}

func Print(args ...interface{}) {
	L.Print(args...)
}

func Warn(args ...interface{}) {
	L.Warn(args)
}

func Error(args ...interface{}) {
	L.Error(args)
}

func Fatal(args ...interface{}) {
	L.Fatal(args)
}

func Panic(args ...interface{}) {
	L.Panic(args)
}

func Debugln(args ...interface{}) {
	L.Debugln(args)
}

func Infoln(args ...interface{}) {
	L.Infoln(args)
}

func Println(args ...interface{}) {
	L.Println(args...)
}

func Warnln(args ...interface{}) {
	L.Warnln(args)
}

func Errorln(args ...interface{}) {
	L.Errorln(args)
}

func Fatalln(args ...interface{}) {
	L.Fatalln(args)
}

func Panicln(args ...interface{}) {
	L.Panicln(args)
}

func AddPayload(addPayload AddPayloadHandler) {
	L.AddPayload(addPayload)
}
