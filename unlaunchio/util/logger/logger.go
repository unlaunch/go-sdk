package logger

import (
	"fmt"
	"log"
)

// Options ...
type Options struct {
	BaseLogger *log.Logger
	Level string
}

// Logger ...
type Logger struct {
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	traceLogger   *log.Logger
}

// Trace ...
func (l *Logger) Trace(msg ...interface{}) {
	l.traceLogger.Println(msg...)
}

// Debug ...
func (l *Logger) Debug(msg ...interface{}) {
	l.debugLogger.Println(msg...)
}

// Info ...
func (l *Logger) Info(msg ...interface{}) {
	l.infoLogger.Println(msg...)
}

// Warning ...
func (l *Logger) Warning(msg ...interface{}) {
	l.warningLogger.Println(msg...)
}

// Error ...
func (l *Logger) Error(msg ...interface{}) {
	l.errorLogger.Println(msg...)
}

// NewLogger ...
func NewLogger(opt *Options) *LevelsLogger {

	lvl := Level(opt.Level)

	opt.BaseLogger.Println()

	opt.BaseLogger.SetPrefix(fmt.Sprintf("%s - unlaunch - ", opt.Level))
	l := &Logger{
		debugLogger:   opt.BaseLogger,
		infoLogger:    opt.BaseLogger,
		warningLogger: opt.BaseLogger,
		errorLogger:   opt.BaseLogger,
		traceLogger:   opt.BaseLogger,
	}

	return &LevelsLogger{
		delegate: l,
		level: lvl,
	}
}
