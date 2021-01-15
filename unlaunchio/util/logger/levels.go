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

	// LevelVerbose log level
	LevelTrace
)


type LevelsLogger struct {
	level    int
	delegate LoggerInterface
}

// Error forwards error logging messages
func (l *LevelsLogger) Error(is ...interface{}) {
	if l.level >= LevelError {
		l.delegate.Error(is...)
	}
}

// Warning forwards warning logging messages
func (l *LevelsLogger) Warning(is ...interface{}) {
	if l.level >= LevelWarning {
		l.delegate.Warning(is...)
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

// Verbose forwards verbose logging messages
func (l *LevelsLogger) Trace(is ...interface{}) {
	if l.level >= LevelTrace {
		l.delegate.Trace(is...)
	}
}

var levels map[string]int = map[string]int{
	"ERROR":   LevelError,
	"WARNING": LevelWarning,
	"INFO":    LevelInfo,
	"DEBUG":   LevelDebug,
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