package tlogger

import (
	"github.com/charmbracelet/log"
	"testing"
)

func TestNewTLogger(t *testing.T) {
	logger := NewTLogger(t)
	logger.Info("hello world")
	logger.Debug("hello world")
}

func TestNewTLoggerWithOptions(t *testing.T) {
	logger := NewTLoggerWithOptions(t, log.Options{
		Level:           log.DebugLevel,
		Prefix:          "Test Prefix",
		ReportTimestamp: true,
		ReportCaller:    true,
	})
	logger.Info("hello world")
	logger.Debug("hello world")
}
