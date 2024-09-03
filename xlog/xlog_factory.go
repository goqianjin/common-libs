package xlog

import (
	"context"
	"io"

	"github.com/goqianjin/common-libs/xlog/internal"
	"github.com/goqianjin/common-libs/xlog/slog"
)

func New(opts ...Option) (Logger, context.Context) {
	opt := &option{}
	for _, o := range opts {
		o(opt)
	}

	// default options
	if opt.Basis == "" {
		opt.Basis = defaultBasis
	}
	if opt.W == nil {
		opt.W = defaultOutput
	}
	if opt.Format == nil {
		opt.Format = &defaultFormat
	}
	if opt.Level == nil {
		opt.Level = &defaultLevel
	}
	if opt.Ctx == nil {
		opt.Ctx = context.Background()
	}
	if opt.AddSource == nil {
		opt.AddSource = defaultAddSource
	}

	// context: requestID
	requestID := opt.requestID
	if requestID == "" && opt.authRequestID != nil && *opt.authRequestID {
		requestID = GenReqID()
	}
	if requestID != "" {
		opt.Ctx = SetAttributeToContext(opt.Ctx, NewAttr("requestID", requestID))
	}

	var basisLogger internal.LoggerInternal
	switch opt.Basis {
	case BasisSlog:
		basisLogger = slog.New(opt.Option, nil)
	}
	l := &loggerAdapter{delegate: basisLogger, ctx: opt.Ctx}

	// inject logger into ctx: to make the package delegate method to work
	ctx := context.WithValue(opt.Ctx, ctxKeyXLogger{}, l)

	return l, ctx
}

func NewFmtLogger(opts ...Option) (FmtLogger, context.Context) {
	l, ctx := New(opts...)
	return l.(*loggerAdapter), ctx
}

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

// ------ options ------

type option struct {
	internal.Option

	Basis Basis // logging basis

	requestID     string // 非空时,优先级高于 AutoRequestID
	authRequestID *bool  // 禁用自动注入 RequestID

}

type Option func(o *option)

func WithContext(ctx context.Context) Option {
	return func(o *option) {
		o.Ctx = ctx
	}
}

func WithRequestID(requestID string) Option {
	return func(o *option) {
		o.requestID = requestID
	}
}

func WithAutoRequestID(autoReqID bool) Option {
	return func(o *option) {
		o.authRequestID = &autoReqID
	}
}

func WithBasis(b Basis) Option {
	return func(o *option) {
		o.Basis = b
	}
}

func WithOutput(w io.Writer) Option {
	return func(o *option) {
		o.W = w
	}
}

func WithFormat(f internal.Format) Option {
	return func(o *option) {
		o.Format = &f
	}
}

func WithLevel(level internal.Level) Option {
	return func(o *option) {
		o.Level = &level
	}
}

func WithSource(addSource bool) Option {
	return func(o *option) {
		o.AddSource = &addSource
	}
}

// ------ wrapper ------

// WrapLogger makes l the default [Logger], which is used by
// the top-level functions [Info], [Debug] and so on.
func WrapLogger(ctx context.Context, basisLogger internal.LoggerInternal) (Logger, context.Context) {
	l := &loggerAdapter{ctx: context.Background(), delegate: basisLogger}
	ctx = context.WithValue(ctx, ctxKeyXLogger{}, l)
	return l, ctx
}

func WrapFmtLogger(ctx context.Context, basisLogger internal.LoggerInternal) (FmtLogger, context.Context) {
	l, ctx := WrapLogger(ctx, basisLogger)
	return l.(*loggerAdapter), ctx
}
