package logger

// LoggerInterface ...
type LoggerInterface interface {
	Trace(msg ...interface{})
	Debug(msg ...interface{})
	Info(msg ...interface{})
	Warning(msg ...interface{})
	Error(msg ...interface{})
}
