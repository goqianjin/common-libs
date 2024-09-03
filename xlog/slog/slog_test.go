package slog

import (
	"log/slog"
	"testing"
)

func TestGoSlog(t *testing.T) {
	slog.Info("hello info1")
	slog.Info("hello info2")
}
