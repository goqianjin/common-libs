package xlog

import (
	"context"
	"fmt"

	"github.com/goqianjin/common-libs/xlog/internal"
)

// ------ interface ------

type Logger interface {
	Trace(msg string, args ...any)
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Fatal(msg string, args ...any)
}

type Logf interface {
	Tracef(format string, args ...any)
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
}

// ------ package logging functions ------

func Trace(ctx context.Context, msg string, args ...any) {
	getLogger(ctx).(*loggerAdapter).Log(LevelTrace, msg, args...)
}

func Debug(ctx context.Context, msg string, args ...any) {
	getLogger(ctx).(*loggerAdapter).Log(LevelDebug, msg, args...)
}
func Info(ctx context.Context, msg string, args ...any) {
	getLogger(ctx).(*loggerAdapter).Log(LevelInfo, msg, args...)
}

func Warn(ctx context.Context, msg string, args ...any) {
	getLogger(ctx).(*loggerAdapter).Log(LevelWarn, msg, args...)
}

func Error(ctx context.Context, msg string, args ...any) {
	getLogger(ctx).(*loggerAdapter).Log(LevelError, msg, args...)
}

func Fatal(ctx context.Context, msg string, args ...any) {
	getLogger(ctx).(*loggerAdapter).Log(LevelFatal, msg, args...)
}

// ------ package formatted logging functions ------

func Tracef(ctx context.Context, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	getLogger(ctx).(*loggerAdapter).Log(LevelTrace, msg)
}

func Debugf(ctx context.Context, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	getLogger(ctx).(*loggerAdapter).Log(LevelDebug, msg)
}

func Infof(ctx context.Context, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	getLogger(ctx).(*loggerAdapter).Log(LevelInfo, msg)
}

func Warnf(ctx context.Context, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	getLogger(ctx).(*loggerAdapter).Log(LevelWarn, msg)
}

func Errorf(ctx context.Context, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	getLogger(ctx).(*loggerAdapter).Log(LevelError, msg)
}

func Fatalf(ctx context.Context, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	getLogger(ctx).(*loggerAdapter).Log(LevelFatal, msg)
}

// ------ Format ------

const (
	FormatJSON      = internal.FormatJSON
	FormatText      = internal.FormatText
	FormatClassical = internal.FormatClassical
)

// ------ Level ------

const (
	LevelTrace = internal.LevelTrace
	LevelDebug = internal.LevelDebug
	LevelInfo  = internal.LevelInfo
	LevelWarn  = internal.LevelWarn
	LevelError = internal.LevelError
	LevelFatal = internal.LevelFatal
)

// ------ BaseOn ------

/*type Basis string

const (
	BasisSlog = Basis("Slog")
)*/

// ------ helpers ------

var (
	PutContextAttribute  = internal.PutContextAttribute
	GetContextAttribute  = internal.GetContextAttribute
	GetContextAttributes = internal.GetContextAttributes
	GetContextReqID      = internal.GetContextReqID
)
