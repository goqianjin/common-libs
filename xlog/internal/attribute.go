package internal

import "log/slog"

type Attr slog.Attr

func NewAttr(key string, v any) Attr {
	return Attr(slog.Any(key, v))
}
