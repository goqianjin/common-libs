package xlog

import (
	"context"
	"fmt"
	"os"

	"github.com/goqianjin/common-libs/xlog/internal"
)

// ------ logger wrapper ------

type loggerAdapter struct {
	delegate internal.LoggerInternal
	ctx      context.Context
}

func (l *loggerAdapter) Log(level internal.Level, msg string, args ...any) {
	l.delegate.Log(l.ctx, level, msg, args...)
	if level == LevelFatal {
		os.Exit(1)
	}
}

// ------ Logger implementation ------

func (l *loggerAdapter) Trace(msg string, args ...any) {
	l.Log(LevelTrace, msg, args...)
}

func (l *loggerAdapter) Debug(msg string, args ...any) {
	l.Log(LevelDebug, msg, args...)
}

func (l *loggerAdapter) Info(msg string, args ...any) {
	l.Log(LevelInfo, msg, args...)
}

func (l *loggerAdapter) Warn(msg string, args ...any) {
	l.Log(LevelWarn, msg, args...)
}

func (l *loggerAdapter) Error(msg string, args ...any) {
	l.Log(LevelError, msg, args...)
}

func (l *loggerAdapter) Fatal(msg string, args ...any) {
	l.Log(LevelFatal, msg, args...)
}

// ------ Logf implementation ------

func (l *loggerAdapter) Tracef(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Log(LevelTrace, msg, args...)
}

func (l *loggerAdapter) Debugf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Log(LevelDebug, msg, args...)
}

func (l *loggerAdapter) Infof(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Log(LevelInfo, msg, args...)
}

func (l *loggerAdapter) Warnf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Log(LevelWarn, msg, args...)
}

func (l *loggerAdapter) Errorf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Log(LevelError, msg, args...)
}

func (l *loggerAdapter) Fatalf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Log(LevelFatal, msg, args...)
}
