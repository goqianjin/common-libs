package xlog

import (
	"io"
	"os"
	"sync/atomic"

	"github.com/goqianjin/common-libs/xlog/internal"
)

// 默认级别与 slog 一致
var defaultLogger atomic.Value

// Default returns the default [Logger].
func Default() Logger {
	l := defaultLogger.Load()
	if l == nil {
		l = newDefaultLogger()
		defaultLogger.Store(l)
	}
	return l.(Logger)
}

func newDefaultLogger() Logger {
	l, _ := New(WithOutput(defaultOutput), WithLevel(defaultLevel),
		WithSource(*defaultAddSource), WithAutoRequestID(*defaultAutoRequestID),
	)
	return l
}

// DefaultFmtLogger returns the default [FmtLogger].
func DefaultFmtLogger() FmtLogger {
	return Default().(*loggerAdapter)
}

var defaultOutput io.Writer = os.Stdout

func SetOutput(w io.Writer) {
	defaultOutput = w
	if defaultLogger.Load() != nil {
		defaultLogger.Store(newDefaultLogger())
	}
}

var defaultLevel internal.Level = LevelDebug

func SetLevel(level internal.Level) {
	defaultLevel = level
	if defaultLogger.Load() != nil {
		defaultLogger.Store(newDefaultLogger())
	}
}

var defaultBasis Basis = BasisSlog

func SetBasis(basis Basis) {
	defaultBasis = basis
	if defaultLogger.Load() != nil {
		defaultLogger.Store(newDefaultLogger())
	}
}

var defaultFormat internal.Format = FormatClassical

func SetFormat(format internal.Format) {
	defaultFormat = format
	if defaultLogger.Load() != nil {
		defaultLogger.Store(newDefaultLogger())
	}
}

var defaultAddSource = func() *bool { addSource := true; return &addSource }()

func SetAddSource(addSource bool) {
	defaultAddSource = &addSource
	if defaultLogger.Load() != nil {
		defaultLogger.Store(newDefaultLogger())
	}
}

var defaultAutoRequestID = func() *bool { addSource := true; return &addSource }()

func WithAutoReqID(autoReqID bool) {
	defaultAutoRequestID = &autoReqID
	if defaultLogger.Load() != nil {
		defaultLogger.Store(newDefaultLogger())
	}
}
