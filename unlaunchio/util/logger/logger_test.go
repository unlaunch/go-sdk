package logger

import (
	"testing"
)

type mockLogger struct {
	msgs map[string]bool
}

func (l *mockLogger) reset() {
	l.msgs = make(map[string]bool)
}

func (l *mockLogger) Error(msg ...interface{}) {
	l.msgs["ERROR"] = true
}

func (l *mockLogger) Warn(msg ...interface{}) {
	l.msgs["WARN"] = true
}

func (l *mockLogger) Info(msg ...interface{}) {
	l.msgs["INFO"] = true
}

func (l *mockLogger) Debug(msg ...interface{}) {
	l.msgs["DEBUG"] = true
}

func (l *mockLogger) Trace(msg ...interface{}) {
	l.msgs["TRACE"] = true
}

func logAtAllLevels(logger *LevelsLogger) {
	logger.Error("And finally…")
	logger.Warn("And finally…")
	logger.Info("And finally…")
	logger.Debug("And finally…")
	logger.Trace("And finally…")
}

func checkShouldBeCalled(delegate *mockLogger, shouldBeCalled []string, t *testing.T) {
	for _, level := range shouldBeCalled {
		if !delegate.msgs[level] {
			t.Errorf("Call to log level function \"%s\" should have been forwarded", level)
		}
	}
}

func checkShouldNotBeCalled(delegate *mockLogger, shouldBeCalled []string, t *testing.T) {
	for _, level := range shouldBeCalled {
		if delegate.msgs[level] {
			t.Errorf("Call to log level function \"%s\" should have been forwarded", level)
		}
	}
}

func TestErrorLevel(t *testing.T) {
	delegate := &mockLogger{}
	delegate.reset()

	logger := LevelsLogger{
		delegate: delegate,
		level: Level("ERROR"),
	}

	logAtAllLevels(&logger)

	shouldBeCalled := []string{"ERROR"}
	shouldNotBeCalled := []string{"WARN", "INFO", "DEBUG", "TRACE"}

	checkShouldBeCalled(delegate, shouldBeCalled, t)
	checkShouldNotBeCalled(delegate, shouldNotBeCalled, t)
}

func TestWarnLevel(t *testing.T) {
	delegate := &mockLogger{}
	delegate.reset()

	logger := LevelsLogger{
		delegate: delegate,
		level: Level("WARN"),
	}

	logAtAllLevels(&logger)

	shouldBeCalled := []string{"ERROR", "WARN"}
	shouldNotBeCalled := []string{"INFO", "DEBUG", "TRACE"}

	checkShouldBeCalled(delegate, shouldBeCalled, t)
	checkShouldNotBeCalled(delegate, shouldNotBeCalled, t)
}

func TestInfoLevel(t *testing.T) {
	delegate := &mockLogger{}
	delegate.reset()

	logger := LevelsLogger{
		delegate: delegate,
		level: Level("INFO"),
	}

	logAtAllLevels(&logger)

	shouldBeCalled := []string{"ERROR", "WARN", "INFO"}
	shouldNotBeCalled := []string{"DEBUG", "TRACE"}

	checkShouldBeCalled(delegate, shouldBeCalled, t)
	checkShouldNotBeCalled(delegate, shouldNotBeCalled, t)
}

func TestDebugLevel(t *testing.T) {
	delegate := &mockLogger{}
	delegate.reset()

	logger := LevelsLogger{
		delegate: delegate,
		level: Level("DEBUG"),
	}

	logAtAllLevels(&logger)

	shouldBeCalled := []string{"DEBUG", "ERROR", "WARN", "INFO"}
	shouldNotBeCalled := []string{"TRACE"}

	checkShouldBeCalled(delegate, shouldBeCalled, t)
	checkShouldNotBeCalled(delegate, shouldNotBeCalled, t)
}

func TestTraceLevel(t *testing.T) {
	delegate := &mockLogger{}
	delegate.reset()

	logger := LevelsLogger{
		delegate: delegate,
		level: Level("TRACE"),
	}

	logAtAllLevels(&logger)

	shouldBeCalled := []string{"TRACE", "DEBUG", "ERROR", "WARN", "INFO"}
	shouldNotBeCalled := []string{}

	checkShouldBeCalled(delegate, shouldBeCalled, t)
	checkShouldNotBeCalled(delegate, shouldNotBeCalled, t)
}

