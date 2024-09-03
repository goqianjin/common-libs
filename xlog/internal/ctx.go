package internal

import "context"

type ctxKeyContextAttributes struct{}

// PutContextAttribute sets a slog attribute to the provided context so that it will be
// included in any Record created with such context.
// the return context will be a renewed one based-on the parent when the first attribute
// is added into the context; otherwise, it's the same reference of the parent.
func PutContextAttribute(parent context.Context, key string, value any) context.Context {
	if parent == nil {
		parent = context.Background()
	}
	ctx := parent

	attrs := GetContextAttributes(ctx)
	if attrs == nil {
		ctx = InitContextAttributes(ctx)
		attrs = GetContextAttributes(ctx)
	}
	attrs[key] = value

	return ctx
}

func GetContextAttribute(ctx context.Context, key string) any {
	attrs := GetContextAttributes(ctx)
	if attrs == nil {
		return nil
	}
	return attrs[key]
}

func GetContextAttributes(ctx context.Context) map[string]any {
	v := ctx.Value(ctxKeyContextAttributes{})
	if v == nil {
		return nil
	}
	if attrs, ok := v.(map[string]any); ok {
		return attrs
	}
	return nil
}

func InitContextAttributes(parent context.Context) context.Context {
	if parent == nil {
		parent = context.Background()
	}
	ctx := parent

	// ctx attrs
	v := ctx.Value(ctxKeyContextAttributes{})
	if v == nil {
		attrs := make(map[string]any)
		ctx = context.WithValue(parent, ctxKeyContextAttributes{}, attrs)
	}

	return ctx
}
