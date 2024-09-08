package rlog

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"sync"
	"time"

	"github.com/goqianjin/common-libs/xlog/internal"
	"github.com/goqianjin/common-libs/xlog/internal/buffer"
)

const (
	FieldNameMessage   = "MSG"
	FieldNameArguments = "ARGS..."

	TimeFormat = "2006-01-02 15:04:05.000000"
)

type rawLogger struct {
	mu *sync.Mutex
	w  io.Writer

	level      internal.Level
	separator  string
	fieldNames []string
}

func (l *rawLogger) Log(ctx context.Context, level internal.Level, msg string, args ...any) {
	if level < internal.DefaultLevel {
		return
	}
	buf := buffer.New()
	defer buf.Free()

	attrs := internal.GetContextAttributes(ctx)

	if len(l.fieldNames) > 0 {
		for _, name := range l.fieldNames {
			switch name {
			case FieldNameMessage:
				_, _ = buf.WriteString(msg)
			case FieldNameArguments:
				for _, arg := range args {
					appendTextValue(buf, arg)
					_, _ = buf.WriteString(l.separator)
				}
			default:
				appendTextValue(buf, attrs[name])
				_, _ = buf.WriteString(l.separator)
			}
		}
	} else {
		// attributes
		for k, v := range attrs {
			_, _ = buf.WriteString(k)
			_ = buf.WriteByte('=')
			appendTextValue(buf, v)
			_, _ = buf.WriteString(l.separator)
		}
		// message
		//_, _ = buf.WriteString(msg)
		appendTextValue(buf, msg)
		_, _ = buf.WriteString(l.separator)
		// arguments
		for _, arg := range args {
			appendTextValue(buf, arg)
			_, _ = buf.WriteString(l.separator)
		}
	}

	// try append '\n' in the end
	if (*buf)[len(*buf)-1] != '\n' {
		_ = buf.WriteByte('\n')
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	_, _ = l.w.Write(*buf)

	return
}

const (
	nilValue = '-'
)

func appendTextValue(buf *buffer.Buffer, v any) {
	if v == nil {
		v = nilValue
	}
	switch v.(type) {
	case time.Time:
		buf.WriteString(strconv.Quote(v.(time.Time).Format(TimeFormat)))
	case string, fmt.Stringer:
		buf.WriteString(strconv.Quote(fmt.Sprint(v)))
	default:
		buf.WriteString(fmt.Sprint(v))
	}
}
