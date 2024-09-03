package slog

import (
	"context"
	"os"
	"testing"

	"github.com/goqianjin/common-libs/xlog/internal"
	"github.com/goqianjin/common-libs/xlog/internal/assert"
)

func TestCtxHandler(t *testing.T) {
	l := New(os.Stdout, Option{})

	ctx := internal.PutContextAttribute(context.Background(), "IntKey", 123)
	ctx = internal.PutContextAttribute(ctx, "StringKey", "hello")
	assert.Equal(t, 2, len(internal.GetContextAttributes(ctx)))

	l.Log(ctx, internal.LevelInfo, "Hello INFO.", "key", "value")
	l.Log(ctx, internal.LevelWarn, "Hello WARN.", "key", "value")
}
