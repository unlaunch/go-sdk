package logger

const (
	_ = iota

	// LevelError log level
	LevelError

	// LevelWarning log level
	LevelWarning

	// LevelInfo log level
	LevelInfo

	// LevelDebug log level
	LevelDebug

	// LevelTrace log level
	LevelTrace
)

// LevelsLogger encapsulates logger and applies log levels
type LevelsLogger struct {
	level    int
	delegate Interface
}

// Error forwards error logging messages
func (l *LevelsLogger) Error(is ...interface{}) {
	if l.level >= LevelError {
		l.delegate.Error(is...)
	}
}

// Warn forwards warning logging messages
func (l *LevelsLogger) Warn(is ...interface{}) {
	if l.level >= LevelWarning {
		l.delegate.Warn(is...)
	}
}

// Info forwards info logging messages
func (l *LevelsLogger) Info(is ...interface{}) {
	if l.level >= LevelInfo {
		l.delegate.Info(is...)
	}
}

// Debug forwards debug logging messages
func (l *LevelsLogger) Debug(is ...interface{}) {
	if l.level >= LevelDebug {
		l.delegate.Debug(is...)
	}
}

// Trace forwards verbose logging messages
func (l *LevelsLogger) Trace(is ...interface{}) {
	if l.level >= LevelTrace {
		l.delegate.Trace(is...)
	}
}

var levels map[string]int = map[string]int{
	"ERROR": LevelError,
	"WARN":  LevelWarning,
	"INFO":  LevelInfo,
	"DEBUG": LevelDebug,
	"TRACE": LevelTrace,
}

// Level gets current level
func Level(level string) int {
	l, ok := levels[level]
	if !ok {
		panic("Invalid log level " + level)
	}
	return l
}
