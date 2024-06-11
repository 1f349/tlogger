package tlogger

import (
	"github.com/charmbracelet/log"
	"strings"
	"testing"
)

type tLogger struct{ t *testing.T }

// NewTLogger creates a new tLogger.
func NewTLogger(t *testing.T) *log.Logger {
	return log.New(&tLogger{t: t})
}

func NewTLoggerWithOptions(t *testing.T, o log.Options) *log.Logger {
	return log.NewWithOptions(&tLogger{t: t}, o)
}

// Logf logs a formatted message to the testing framework.
func (tl *tLogger) Logf(format string, args ...interface{}) {
	tl.t.Logf(format, args...)
}

// Write implements the io.Writer interface.
func (tl *tLogger) Write(p []byte) (n int, err error) {
	tl.t.Log(strings.TrimSpace(string(p)))
	return len(p), nil
}
