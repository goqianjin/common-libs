package slog

import (
	"context"
	"io"
	"log/slog"
	"runtime"
	"strconv"
	"sync"

	"github.com/goqianjin/common-libs/xlog/internal"
	"github.com/goqianjin/common-libs/xlog/internal/buffer"
)

var logTimeFormat = "2006/01/02 15:04:05.000000"

// Reference:
// (1) https://github.com/golang/go/blob/4f852b9734249c063928b34a02dd689e03a8ab2c/src/log/slog/text_handler.go#L94-L96
// (2) https://github.com/golang/go/blob/4f852b9734249c063928b34a02dd689e03a8ab2c/src/log/slog/handler.go#L265-L315

// classicalHandler implements slog.Handler in a classical logging formatter.
// The formatter like: {time}: TODO complete
type classicalHandler struct {
	mu *sync.Mutex
	w  io.Writer

	level     *internal.Level
	addSource *bool

	// internal fixed field
	attrPrefix string
}

func newClassicalHandler(w io.Writer, opts *slog.HandlerOptions) *classicalHandler {
	level := internal.DefaultLevel
	if opts.Level != nil {
		level = opts.Level.Level()
	}

	mutex := &sync.Mutex{}
	if w == internal.DefaultW { // 是默认输出, 使用默认Lock
		mutex = &internal.DefaultWMutex
	}
	h := &classicalHandler{
		mu:        mutex,
		w:         w,
		level:     &level,
		addSource: &opts.AddSource,
	}

	return h
}

func (h *classicalHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := buffer.New()
	defer buf.Free()

	// time
	if !r.Time.IsZero() {
		*buf = r.Time.AppendFormat(*buf, logTimeFormat)
	}
	_ = buf.WriteByte(' ') // separator

	// custom - requestID
	if reqID := internal.GetContextReqID(ctx); reqID != "" {
		_ = buf.WriteByte('[')
		_, _ = buf.WriteString(reqID)
		_ = buf.WriteByte(']')
	}

	// level
	_ = buf.WriteByte('[')
	_, _ = buf.WriteString(internal.FormatLevelToString(r.Level))
	_ = buf.WriteByte(']')
	_ = buf.WriteByte(' ') // separator

	// source
	if h.addSource != nil && *h.addSource {
		source := getSourceBySlogRecord(r)
		source.File = internal.FormatFile(source.File, 2)
		_, _ = buf.WriteString(source.File)
		_ = buf.WriteByte(':')
		_, _ = buf.WriteString(strconv.Itoa(source.Line))
		_ = buf.WriteByte(':')
		_ = buf.WriteByte(' ') // separator
	}

	// extra attr
	if h.attrPrefix != "" {
		_, _ = buf.WriteString(h.attrPrefix)
		_ = buf.WriteByte(' ') // separator
	}
	if r.NumAttrs() > 0 {
		// logging attributes
		r.Attrs(func(attr slog.Attr) bool {
			// skip previous logging-head attributes
			if attr.Key == slog.SourceKey || attr.Key == internal.AttributeKeyReqID {
				return true
			}
			_, _ = buf.WriteString(attr.Key)
			_ = buf.WriteByte('=')
			appendTextValue(buf, attr.Value)
			_ = buf.WriteByte(' ') // separator
			return true
		})
	}

	_, _ = buf.WriteString(r.Message)
	if r.Message[len(r.Message)-1] != '\n' {
		_ = buf.WriteByte('\n')
	}

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.w.Write(*buf)

	return err
}

func (h *classicalHandler) Enabled(ctx context.Context, level slog.Level) bool {
	if h.level != nil {
		return level >= *h.level
	}
	return level >= internal.DefaultLevel
}

func (h *classicalHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	handler := h.clone()

	// format attributes
	buf := buffer.New()
	defer buf.Free()
	for _, attr := range attrs {
		_, _ = buf.WriteString(attr.Key)
		_ = buf.WriteByte('=')
		appendTextValue(buf, attr.Value)
		_ = buf.WriteByte(' ')
	}
	handler.attrPrefix = buf.String()

	return handler
}

func (h *classicalHandler) WithGroup(name string) slog.Handler {
	panic("not implement")
}

func (h *classicalHandler) clone() *classicalHandler {
	// We can't use assignment because we can't copy the mutex.
	return &classicalHandler{
		w:  h.w,
		mu: h.mu, // mutex shared among all clones of this handler

		level:     h.level,
		addSource: h.addSource,
	}
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

func appendTextValue(buf *buffer.Buffer, v slog.Value) {
	switch v.Kind() {
	case slog.KindString, slog.KindTime:
		buf.WriteString(strconv.Quote(v.String()))
	default:
		buf.WriteString(v.String())
	}
}
