package slog

import (
	"io"
	"log/slog"

	"github.com/goqianjin/common-libs/xlog/internal"
)

func NewTextHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	rewriteReplaceAttr(opts)
	return &handlerAdapter{slog.NewTextHandler(w, opts)}
}

func NewJSONHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	rewriteReplaceAttr(opts)
	return &handlerAdapter{slog.NewJSONHandler(w, opts)}
}

func NewPrettyHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	rewriteReplaceAttr(opts)
	return &handlerAdapter{newClassicalHandler(w, opts)}
}

func rewriteReplaceAttr(opts *slog.HandlerOptions) {
	if opts == nil {
		return
	}
	if opts.ReplaceAttr != nil {
		opts.ReplaceAttr = func(groups []string, a slog.Attr) slog.Attr {
			a = opts.ReplaceAttr(groups, a)
			a = replaceLevelAttr(groups, a)
			a = replaceSourceAttr(groups, a)

			return a
		}
	} else {
		opts.ReplaceAttr = func(groups []string, a slog.Attr) slog.Attr {
			a = replaceLevelAttr(groups, a)
			a = replaceSourceAttr(groups, a)
			return a
		}
	}
}

func replaceLevelAttr(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		v := a.Value.Any().(slog.Level)
		vLabel := v.String()

		switch v {
		case internal.LevelTrace:
			// NOTE: 如果不设置，默认日志级别打印为 "level":"DEBUG+2"
			vLabel = "TRACE"
		case internal.LevelFatal:
			vLabel = "FATAL"
		}
		a.Value = slog.StringValue(vLabel)
	}
	return a
}

func replaceSourceAttr(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.SourceKey {
		v := a.Value.Any().(*slog.Source)
		v.File = internal.FormatFile(v.File, 2)
	}
	return a
}
