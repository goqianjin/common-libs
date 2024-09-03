package xlog

import (
	"context"
	"io"

	"github.com/goqianjin/common-libs/xlog/internal"
	"github.com/goqianjin/common-libs/xlog/rlog"
	"github.com/goqianjin/common-libs/xlog/slog"
)

func NewSLog(w io.Writer, opt SlogOption) (Logger, context.Context) {
	if opt.Ctx == nil {
		opt.Ctx = context.Background()
	}
	if opt.AddSource == nil {
		opt.AddSource = &internal.DefaultAddSource
	}
	if opt.AutoReqID == nil {
		opt.AutoReqID = &internal.DefaultAutoReqID
	}

	// context: in order to reserve attributes in a fixed contain
	opt.Ctx = internal.InitContextAttributes(opt.Ctx)

	// context: reqID
	requestID := opt.ReqID
	if requestID == "" && *opt.AutoReqID {
		requestID = internal.GenReqID()
	}
	if requestID != "" {
		opt.Ctx = PutContextAttribute(opt.Ctx, internal.AttributeKeyReqID, requestID)
	}

	// creating logger
	basisLogger := slog.New(w, slog.Option{
		Format:    opt.Format,
		Level:     opt.Level,
		AddSource: opt.AddSource,
		Args:      opt.Args,
	})
	l := &loggerAdapter{delegate: basisLogger, ctx: opt.Ctx}

	// inject logger into ctx: to make the package delegate method to work
	ctx := context.WithValue(opt.Ctx, ctxKeyXLogger{}, l)

	return l, ctx
}

func NewSLogf(w io.Writer, opt SlogOption) (Logf, context.Context) {
	l, ctx := NewSLog(w, opt)
	return l.(*loggerAdapter), ctx
}

func NewRLog(w io.Writer, opt RawLogOption) (Logger, context.Context) {
	if opt.Ctx == nil {
		opt.Ctx = context.Background()
	}
	if opt.AutoReqID == nil {
		opt.AutoReqID = &internal.DefaultAutoReqID
	}

	// context: in order to reserve attributes in a fixed contain
	opt.Ctx = internal.InitContextAttributes(opt.Ctx)

	// context: reqID
	requestID := opt.ReqID
	if requestID == "" && *opt.AutoReqID {
		requestID = internal.GenReqID()
	}
	if requestID != "" {
		opt.Ctx = PutContextAttribute(opt.Ctx, internal.AttributeKeyReqID, requestID)
	}

	// creating logger
	basisLogger := rlog.New(w, rlog.Option{
		Level: opt.Level,
	})
	l := &loggerAdapter{delegate: basisLogger, ctx: opt.Ctx}

	// inject logger into ctx: to make the package delegate method to work
	ctx := context.WithValue(opt.Ctx, ctxKeyXLogger{}, l)

	return l, ctx
}

func NewRLogf(w io.Writer, opt RawLogOption) (Logf, context.Context) {
	l, ctx := NewRLog(w, opt)
	return l.(*loggerAdapter), ctx
}

func New(options ...Option) (Logger, context.Context) {
	opt := &option{}
	for _, o := range options {
		o(opt)
	}
	if opt.Output == nil {
		opt.Output = internal.DefaultW
	}
	if opt.Level == nil {
		opt.Level = &internal.DefaultLevel
	}
	return NewSLog(opt.Output, SlogOption{Level: opt.Level})
}

// ------ options ------

type SlogOption struct {
	// slog options

	Format    internal.Format // Optional: default is Text
	Level     *internal.Level // Optional: default is INFO
	AddSource *bool           // Optional: default is true
	Args      []any           // Optional:

	// xlog options

	Ctx       context.Context
	ReqID     string
	AutoReqID *bool
}

type RawLogOption struct {
	// rlog options

	Level     *internal.Level // Optional: default is INFO
	Separator string          // Optional: default is '	'

	// xlog options

	Ctx       context.Context
	ReqID     string
	AutoReqID *bool
}

type option struct {
	Output io.Writer
	Level  *internal.Level
}

type Option func(o *option)

func WithOutput(output io.Writer) Option {
	return func(o *option) {
		o.Output = output
	}
}

func WithLevel(level internal.Level) Option {
	return func(o *option) {
		o.Level = &level
	}
}

// ------ context logger ------

type ctxKeyXLogger struct{}

func getLogger(ctx context.Context) Logger {
	xLoggerVal := ctx.Value(ctxKeyXLogger{})
	if xLoggerVal != nil {
		if xLogger, ok := xLoggerVal.(Logger); ok {
			return xLogger
		}
	}
	return Default()
}

// ------ wrapper ------

// WrapLogger makes l the default [Logger], which is used by
// the top-level functions [Info], [Debug] and so on.
func WrapLogger(ctx context.Context, basisLogger internal.LoggerInternal) (Logger, context.Context) {
	l := &loggerAdapter{ctx: context.Background(), delegate: basisLogger}
	ctx = context.WithValue(ctx, ctxKeyXLogger{}, l)
	return l, ctx
}

func WrapLogf(ctx context.Context, basisLogger internal.LoggerInternal) (Logf, context.Context) {
	l, ctx := WrapLogger(ctx, basisLogger)
	return l.(*loggerAdapter), ctx
}
