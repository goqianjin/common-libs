package slog

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"strings"
	"sync"

	"github.com/goqianjin/common-libs/xlog/internal"
	"github.com/goqianjin/common-libs/xlog/internal/buffer"
)

// Reference:
// (1) https://github.com/golang/go/blob/4f852b9734249c063928b34a02dd689e03a8ab2c/src/log/slog/text_handler.go#L94-L96
// (2) https://github.com/golang/go/blob/4f852b9734249c063928b34a02dd689e03a8ab2c/src/log/slog/handler.go#L265-L315

// ClassicalHandler implements slog.Handler in a classical logging formatter.
// The formatter like: {time}: TODO complete
type ClassicalHandler struct {
	*slog.TextHandler

	opts *slog.HandlerOptions
	mu   sync.Mutex
	w    io.Writer
}

func newClassicalHandler(w io.Writer, opts *slog.HandlerOptions) *ClassicalHandler {
	h := &ClassicalHandler{
		w:           w,
		opts:        opts,
		TextHandler: slog.NewTextHandler(w, opts),
	}
	return h
}

var logTimeFormat = "2006/01/02 15:04:05.000000"

func (h *ClassicalHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := buffer.New()
	defer buf.Free()

	// time
	if !r.Time.IsZero() {
		_, _ = buf.WriteString(r.Time.Format(logTimeFormat))
	}
	// custom - requestID
	reqID := ""
	if attrs := internal.GetAttributesFromContext(ctx); len(attrs) > 0 {
		for _, v := range attrs {
			if v.Key == "requestID" {
				reqID = v.Value.String()
				break
			}
		}
	}
	_, _ = buf.WriteString(fmt.Sprintf(" [%v]", reqID))

	// level
	_, _ = buf.WriteString(fmt.Sprintf("[%v]", internal.FormatLevelToString(r.Level)))

	// source
	if h.opts.AddSource {
		source := getSourceBySlogRecord(r)
		source.File = internal.FormatFile(source.File, 2)
		_, _ = buf.WriteString(fmt.Sprintf(" %s:%v:", source.File, source.Line))
	}

	// extra attr
	attrElements := make([]string, 0)
	r.Attrs(func(attr slog.Attr) bool {
		// skip previous logging-head attributes
		if attr.Key == slog.SourceKey || attr.Key == "requestID" {
			return true
		}
		attrElements = append(attrElements, fmt.Sprintf("%s=%s", attr.Key, attr.Value.String()))
		return true
	})
	if len(attrElements) > 0 {
		_, _ = buf.WriteString(fmt.Sprintf(" {%s} --> ", strings.Join(attrElements, ", ")))
	}
	_, _ = buf.WriteString(fmt.Sprintf(" %s", r.Message))
	if r.Message[len(r.Message)-1] != '\n' {
		_ = buf.WriteByte('\n')
	}

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.w.Write(*buf)

	return err
}

// Reference:
// (1) https://github.com/golang/go/blob/239666cd7343d46c40a5b929c8bec8b532dbe83f/src/log/slog/record.go#L214-L226
// (2) https://github.com/golang/go/blob/239666cd7343d46c40a5b929c8bec8b532dbe83f/src/log/slog/logger.go#L245-L251
func getSourceBySlogRecord(r slog.Record) *slog.Source {
	fs := runtime.CallersFrames([]uintptr{r.PC})
	f, _ := fs.Next()
	return &slog.Source{
		Function: f.Function,
		File:     f.File,
		Line:     f.Line,
	}
}
