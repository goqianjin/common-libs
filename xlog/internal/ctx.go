package internal

import "context"

type ctxKeyLogAttributes struct{}

// SetAttributeToContext sets a slog attribute to the provided context so that it will be
// included in any Record created with such context.
// the return context will be a renewed one based-on the parent when the first attribute
// is added into the context; otherwise, it's the same reference of the parent.
func SetAttributeToContext(parent context.Context, attr Attr) context.Context {
	if parent == nil {
		parent = context.Background()
	}
	ctx := parent

	var attrsP *[]Attr
	v := ctx.Value(ctxKeyLogAttributes{})
	if v == nil {
		attrs := make([]Attr, 0)
		attrsP = &attrs
		ctx = context.WithValue(parent, ctxKeyLogAttributes{}, attrsP)
	} else {
		attrsP = v.(*[]Attr) // 内部key, v一定是 *[]Attr 类型, 不需判空
	}

	var updated bool
	for index, e := range *attrsP {
		if e.Key == attr.Key {
			(*attrsP)[index] = attr
			updated = true
			break
		}
	}
	if !updated {
		*attrsP = append(*attrsP, attr)
	}

	return ctx
}

func GetAttributeFromContext(ctx context.Context, key string) any {
	attrs := GetAttributesFromContext(ctx)
	for _, attr := range attrs {
		if attr.Key == key {
			return attr.Value.Any()
		}
	}
	return nil
}

func GetAttributesFromContext(ctx context.Context) []Attr {
	v := ctx.Value(ctxKeyLogAttributes{})
	if v == nil {
		return nil
	}
	if attrs, ok := v.(*[]Attr); ok {
		return *attrs
	}
	return nil
}
