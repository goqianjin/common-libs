package internal

import (
	"context"
	"io"
	"log/slog"
)

// ------ interface ------

type LoggerInternal interface {
	Log(ctx context.Context, level Level, msg string, args ...any)
}

// ------ Level ------

type Level = slog.Level

const (
	LevelTrace = slog.LevelDebug * 2 // 自定义日志级别
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
	LevelPanic = slog.LevelError * 2 // 自定义日志级别
	LevelFatal = slog.LevelError * 4 // 自定义日志级别
)

func FormatLevelToString(l Level) string {
	switch l {
	case LevelTrace:
		return "TRACE"
	case LevelPanic:
		return "PANIC"
	case LevelFatal:
		return "FATAL"
	default:
		return slog.Level(l).String()
	}
}

// ------ Format ------

type Format string

const (
	FormatJSON      Format = "json"
	FormatText      Format = "text"
	FormatClassical Format = "classical"
)

// ------ Option ------

type Opt struct {
	Ctx context.Context

	W         io.Writer // log output, default is os.Stdout
	Format    *Format   //
	Level     *Level    // Handler option:
	AddSource *bool     // Handler option:
	Args      []any     // Optional:
}

type Option func(o *Opt)

// ------ options base ------

func WithContext(ctx context.Context) Option {
	return func(o *Opt) {
		o.Ctx = ctx
	}
}

func WithOutput(w io.Writer) Option {
	return func(o *Opt) {
		o.W = w
	}
}

func WithFormat(f Format) Option {
	return func(o *Opt) {
		o.Format = &f
	}
}

func WithLevel(level Level) Option {
	return func(o *Opt) {
		o.Level = &level
	}
}

func WithAddSource(addSource bool) Option {
	return func(o *Opt) {
		o.AddSource = &addSource
	}
}

func WithArgs(args ...any) Option {
	return func(o *Opt) {
		o.Args = args
	}
}
