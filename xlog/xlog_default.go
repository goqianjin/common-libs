package xlog

import (
	"context"
	"io"
	"sync/atomic"

	"github.com/goqianjin/common-libs/xlog/internal"
)

// 默认级别与 slog 一致
var defaultLogger atomic.Value
var defaultCustomized atomic.Bool

// Default returns the default [Logger].
func Default() Logger {
	l := defaultLogger.Load()
	if l == nil { // lazy load
		l, _ = NewSLog(internal.DefaultW, SlogOption{})
		defaultLogger.Store(l)
	}
	return l.(*loggerAdapter)
}

func newDefaultLogger() Logger {
	l, _ := NewSLog(internal.DefaultW, SlogOption{})
	return l
}

// DefaultLogf returns the default [Logf].
func DefaultLogf() Logf {
	return Default().(*loggerAdapter)
}

func SetDefault(basisLogger internal.LoggerInternal) {
	if basisLogger == nil {
		panic("basisLogger is nil")
	}
	defaultLogger.Store(&loggerAdapter{delegate: basisLogger, ctx: context.Background()})
	defaultCustomized.Store(true)
}

func SetDefaultOutput(w io.Writer) {
	internal.DefaultW = w

	// refresh default logger
	if defaultLogger.Load() != nil && !defaultCustomized.Load() {
		l, _ := NewSLog(internal.DefaultW, SlogOption{})
		defaultLogger.Store(l)
	}
}
