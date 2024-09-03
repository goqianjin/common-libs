package internal

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextAttributes(t *testing.T) {
	parent := context.Background()

	// Set And Get Attribute
	intV := int64(123)
	ctx := SetAttributeToContext(parent, NewAttr("KeyInt", intV))
	assert.Equal(t, intV, GetAttributeFromContext(ctx, "KeyInt"))
	assert.Equal(t, nil, GetAttributeFromContext(parent, "KeyInt"))

	// Set And Get Attribute
	stringV := "stringValue"
	ctx = SetAttributeToContext(ctx, NewAttr("KeyString", stringV))
	assert.Equal(t, stringV, GetAttributeFromContext(ctx, "KeyString"))
	assert.Equal(t, nil, GetAttributeFromContext(parent, "KeyString"))

	// Get Attributes
	assert.Equal(t, 2, len(GetAttributesFromContext(ctx)))
	assert.Equal(t, 0, len(GetAttributesFromContext(parent)))
}
