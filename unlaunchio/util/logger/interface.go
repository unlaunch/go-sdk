package logger

// Interface ...
type Interface interface {
	Trace(msg ...interface{})
	Debug(msg ...interface{})
	Info(msg ...interface{})
	Warning(msg ...interface{})
	Error(msg ...interface{})
}
