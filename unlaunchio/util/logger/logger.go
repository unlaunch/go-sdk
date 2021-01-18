package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Colors
const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

// LogOptions ...
type LogOptions struct {
	ErrorWriter   io.Writer
	WarningWriter io.Writer
	InfoWriter    io.Writer
	DebugWriter   io.Writer
	TraceWriter   io.Writer
	Level         string
	Colorful      bool
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

// Warn ...
func (l *Logger) Warn(msg ...interface{}) {
	l.warningLogger.Println(msg...)
}

// Error ...
func (l *Logger) Error(msg ...interface{}) {
	l.errorLogger.Println(msg...)
}

// NewLogger ...
func NewLogger(opt *LogOptions) *LevelsLogger {
	opt = normalizeOptions(opt)

	var (
		debugPrefix = "DEBUG - unlaunch - "
		infoPrefix  = "INFO - unlaunch - "
		warnPrefix  = "WARN - unlaunch - "
		errorPrefix = "ERROR - unlaunch - "
		tracePrefix = "TRACE - unlaunch - "
	)

	if opt.Colorful {
		debugPrefix = Yellow + "DEBUG" + Reset + " - unlaunch - "
		infoPrefix = Green + "INFO" + Reset + " - unlaunch - "
		warnPrefix = Magenta + "WARN" + Reset + " - unlaunch - "
		errorPrefix = Red + "ERROR" + Reset + " - unlaunch - "
		tracePrefix = Cyan + "TRACE" + Reset + " - unlaunch - "

	}

	flags := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lmsgprefix

	l := &Logger{
		debugLogger:   log.New(opt.DebugWriter, debugPrefix, flags),
		infoLogger:    log.New(opt.InfoWriter, infoPrefix, flags),
		warningLogger: log.New(opt.WarningWriter, warnPrefix, flags),
		errorLogger:   log.New(opt.ErrorWriter, errorPrefix, flags),
		traceLogger:   log.New(opt.TraceWriter, tracePrefix, flags),
	}

	return &LevelsLogger{
		delegate: l,
		level:    Level(opt.Level),
	}
}

func normalizeOptions(opt *LogOptions) *LogOptions {
	var res *LogOptions

	if opt == nil {
		res = &LogOptions{}
		res.Colorful = true
	} else {
		res = opt
	}

	if res.DebugWriter == nil {
		res.DebugWriter = os.Stdout
	}

	if res.InfoWriter == nil {
		res.InfoWriter = os.Stdout
	}

	if res.WarningWriter == nil {
		res.WarningWriter = os.Stdout
	}

	if res.ErrorWriter == nil {
		res.ErrorWriter = os.Stdout
	}

	if res.TraceWriter == nil {
		res.TraceWriter = os.Stdout
	}


	switch res.Level {
	case "ERROR", "WARN", "INFO", "DEBUG", "TRACE":
	default:
		fmt.Println(fmt.Sprintf("Unlaunch - %s is not valid log level. Choose from one of [ERROR, WARN, INFO, DEBUG, TRACE]. Defaulting to ERROR", res.Level))
		res.Level = "ERROR"
	}

	return res
}
