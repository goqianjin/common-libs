package xlog

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequestIDFromContext(t *testing.T) {
	reqID := GenReqID()
	ctx := SetAttributeToContext(context.Background(), NewAttr(attributeKeyRequestID, reqID))
	parsedReqID := GetRequestIDFromContext(ctx)
	assert.Equal(t, reqID, parsedReqID)
}
