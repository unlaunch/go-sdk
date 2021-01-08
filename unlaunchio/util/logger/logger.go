package logger

import (
	"io"
	"log"
	"os"
)

// Options ...
type Options struct {
	CommonWriter io.Writer
	ErrorWriter  io.Writer
}

// Logger ...
type Logger struct {
	debugLogger   log.Logger
	infoLogger    log.Logger
	warningLogger log.Logger
	errorLogger   log.Logger
	traceLogger   log.Logger
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
func NewLogger(options *Options) *Logger {
	var errorWriter io.Writer
	var commonWriter io.Writer


	if options == nil || options.ErrorWriter == nil {
		errorWriter = os.Stdout
	} else {
		errorWriter = options.ErrorWriter
	}

	if options == nil || options.CommonWriter == nil {
		commonWriter = os.Stdout
	} else {
		commonWriter = options.CommonWriter
	}

	return &Logger{
		debugLogger:   *log.New(commonWriter, "DEBUG - ", 1),
		infoLogger:    *log.New(commonWriter, "INFO - ", 1),
		warningLogger: *log.New(commonWriter, "WARNING - ", 1),
		errorLogger:   *log.New(errorWriter, "ERROR - ", 1),
		traceLogger:   *log.New(errorWriter, "TRACE - ", 1),
	}
}
