package slog

import (
	"context"
	"log/slog"
	"runtime"

	"github.com/goqianjin/common-libs/xlog/internal"
)

// handlerAdapter adapts slog.Handler for xlog.
// (1) automatically add context attributes (such as reqID, etc.) into slog.Record. Please note that
// both slog.TextHandler and slog.JSONHandler use context parameter to do nothing.
// (2) adjust caller skip
type handlerAdapter struct {
	slog.Handler
}

// Handle adds contextual attributes to the Record before calling the underlying handler
func (h handlerAdapter) Handle(ctx context.Context, r slog.Record) error {
	// add context attributes
	if attrs := internal.GetAttributesFromContext(ctx); len(attrs) > 0 {
		for _, v := range attrs {
			r.AddAttrs(slog.Attr(v))
		}
	}
	// adjust fixed caller skip
	if r.PC > 0 {
		// SLOG(*slog.Logger.log): skip [runtime.Callers, this function, this function's caller]
		// NOTE: 这里修改 skip 为 6， 源码中 skip 为 3
		// skip: [runtime.Callers, this function, slog.(*Logger).log, slog.(*Logger).Log,
		// xlog.(*loggerAdapter).Log, loggerAdapter.<Function> or xlog package caller]
		var pcs [1]uintptr
		runtime.Callers(6, pcs[:])
		r.PC = pcs[0]
	}

	return h.Handler.Handle(ctx, r)
}
