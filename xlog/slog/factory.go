package slog

import (
	"context"
	"log/slog"

	"github.com/goqianjin/common-libs/xlog/internal"
)

func New(opt internal.Option, handler slog.Handler) internal.LoggerInternal {
	if opt.Ctx == nil {
		opt.Ctx = context.Background()
	}
	if opt.Format == nil {
		format := internal.FormatText
		opt.Format = &format
	}

	handlerOption := &slog.HandlerOptions{
		AddSource: *opt.AddSource,
		Level:     opt.Level,
		//ReplaceAttr: opt.ReplaceAttr,
	}

	if handler == nil {
		switch *opt.Format {
		case internal.FormatJSON:
			handler = NewJSONHandler(opt.W, handlerOption)
		case internal.FormatClassical:
			handler = NewPrettyHandler(opt.W, handlerOption)
		default:
			handler = NewTextHandler(opt.W, handlerOption)
		}
	} else {
		switch handler.(type) {
		case *slog.TextHandler:
			handler = NewTextHandler(opt.W, handlerOption)
		case *slog.JSONHandler:
			handler = NewJSONHandler(opt.W, handlerOption)
		}
	}

	return slog.New(handler)
}
