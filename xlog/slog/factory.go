package slog

import (
	"io"
	"log/slog"

	"github.com/goqianjin/common-libs/xlog/internal"
)

func New(w io.Writer, option Option) internal.LoggerInternal {

	if option.Format == "" {
		option.Format = internal.DefaultFormat
	}
	if option.Level == nil {
		option.Level = &internal.DefaultLevel
	}
	if option.AddSource == nil {
		option.AddSource = &internal.DefaultAddSource
	}

	var handler slog.Handler
	slogOption := &slog.HandlerOptions{AddSource: *option.AddSource, Level: option.Level}
	rewriteReplaceAttr(slogOption)
	switch option.Format {
	case internal.FormatJSON:
		handler = &handlerAdapter{Handler: slog.NewJSONHandler(w, slogOption), RenewPC: *option.AddSource}
	case internal.FormatClassical:
		handler = &handlerAdapter{Handler: newClassicalHandler(w, slogOption), RenewPC: *option.AddSource}
	default:
		handler = &handlerAdapter{Handler: slog.NewTextHandler(w, slogOption), RenewPC: *option.AddSource}
	}
	return slog.New(handler).With(option.Args...)
}

// ------ options base ------

type Option struct {
	Format    internal.Format // Optional: default is Text
	Level     *internal.Level // Optional: default is INFO
	AddSource *bool           // Optional: default is true
	Args      []any           // Optional:
}
