package slog

import (
	"context"
	"log/slog"
	"runtime"
	"sync/atomic"

	"github.com/goqianjin/common-libs/xlog/internal"
)

var renewPC atomic.Bool

// handlerAdapter adapts slog.Handler for xlog.
// (1) automatically add context attributes (such as reqID, etc.) into slog.Record. Please note that
// both slog.TextHandler and slog.JSONHandler use context parameter to do nothing.
// (2) adjust caller skip
type handlerAdapter struct {
	slog.Handler

	RenewPC bool
}

// Handle adds contextual attributes to the Record before calling the underlying handler
func (h handlerAdapter) Handle(ctx context.Context, r slog.Record) error {
	// add context attributes
	if attrs := internal.GetContextAttributes(ctx); len(attrs) > 0 {
		for k, v := range attrs {
			r.AddAttrs(slog.Any(k, v))
		}
	}

	// adjust fixed caller skip
	if renewPC.Load() { // rewrite p.PC
		// SLOG(*slog.Logger.log): skip [runtime.Callers, this function, this function's caller]
		// NOTE: 这里修改 skip 为 6， 源码中 skip 为 3
		// skip: [runtime.Callers, (text/json/classical/custom slog.Handler).Handle, this function,
		//slog.(*Logger).log, slog.(*Logger).Log,
		// xlog.(*loggerAdapter).Log, loggerAdapter.<Function> or xlog package caller]
		var pcs [1]uintptr
		runtime.Callers(6, pcs[:])
		r.PC = pcs[0]
	}

	return h.Handler.Handle(ctx, r)
}
