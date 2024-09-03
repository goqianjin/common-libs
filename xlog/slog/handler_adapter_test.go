package slog

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/goqianjin/common-libs/xlog/internal"
	"github.com/stretchr/testify/assert"
)

func TestCtxHandler(t *testing.T) {
	h := handlerAdapter{
		Handler: slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
		}),
	}
	l := New(internal.Option{}, h)

	ctx := internal.SetAttributeToContext(context.Background(), internal.NewAttr("IntKey", 123))
	ctx = internal.SetAttributeToContext(ctx, internal.NewAttr("StringKey", "hello"))
	assert.Equal(t, 2, len(internal.GetAttributesFromContext(ctx)))

	l.Log(ctx, internal.LevelInfo, "Hello INFO.", slog.Any("key", "value"))
	l.Log(ctx, internal.LevelWarn, "Hello WARN.", slog.Any("key", "value"))
}
