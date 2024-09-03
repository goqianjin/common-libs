package internal

import (
	"context"
	"testing"

	"github.com/goqianjin/common-libs/xlog/internal/assert"
)

func TestGetRequestIDFromContext(t *testing.T) {
	reqID := genReqID()
	ctx := PutContextAttribute(context.Background(), AttributeKeyReqID, reqID)
	parsedReqID := GetContextReqID(ctx)
	assert.Equal(t, reqID, parsedReqID)
}
