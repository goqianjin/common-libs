package slog

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/goqianjin/common-libs/xlog/internal"
)

func TestPrettyHandler(t *testing.T) {
	h := newClassicalHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	})
	l := slog.New(h)
	l.Log(context.Background(), internal.LevelInfo, "Hello INFO.", slog.Any("key", "value"))
	l.Log(context.Background(), internal.LevelWarn, "Hello WARN.", slog.Any("key", "value"))
}
