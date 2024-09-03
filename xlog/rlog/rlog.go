package rlog

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"sync"

	"github.com/goqianjin/common-libs/xlog/internal"
	"github.com/goqianjin/common-libs/xlog/internal/buffer"
)

type rawLogger struct {
	mu *sync.Mutex
	w  io.Writer

	level     internal.Level
	separator string
}

func (l *rawLogger) Log(ctx context.Context, level internal.Level, msg string, args ...any) {
	if level < internal.DefaultLevel {
		return
	}
	buf := buffer.New()
	defer buf.Free()

	attrs := internal.GetContextAttributes(ctx)
	for _, v := range attrs {
		_, _ = buf.WriteString(fmt.Sprint(v))
		_, _ = buf.WriteString(l.separator)
	}

	if len(args) > 0 {
		for _, arg := range args {
			appendTextValue(buf, arg)
			_, _ = buf.WriteString(l.separator)
		}
	}

	_, _ = buf.WriteString(msg)
	if msg[len(msg)-1] != '\n' {
		_ = buf.WriteByte('\n')
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	_, _ = l.w.Write(*buf)

	return
}

type c struct {
}

func (l *c) String() string {
	return ""
}

func appendTextValue(buf *buffer.Buffer, v any) {
	switch v.(type) {
	case string, fmt.Stringer:
		buf.WriteString(strconv.Quote(fmt.Sprint(v)))
	default:
		buf.WriteString(fmt.Sprint(v))
	}
}
