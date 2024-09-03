package xlog

import (
	"context"
	"encoding/base64"
	"encoding/binary"
	"os"
	"time"
)

func GenReqID() string {
	return genReqId()
}

func SetGenReqIDFunc(f func() string) {
	if f == nil {
		f = defaultGenReqId
	}
	genReqId = f
}

// ------ helpers ------
// ReqID implementation
// @see https://github.com/qbox/kodo/blob/develop/libs/xlog.v1/xlog.go#L74-L97

var pid = uint32(os.Getpid())

var genReqId = defaultGenReqId

func defaultGenReqId() string {
	var b [12]byte
	binary.LittleEndian.PutUint32(b[:], pid)
	binary.LittleEndian.PutUint64(b[4:], uint64(time.Now().UnixNano()))
	return base64.URLEncoding.EncodeToString(b[:])
}

// ------ utils ------

var attributeKeyRequestID = "requestID"

func GetRequestIDFromContext(ctx context.Context) string {
	ctxAttrs := GetAttributesFromContext(ctx)
	for _, attr := range ctxAttrs {
		if attr.Key == attributeKeyRequestID {
			return attr.Value.String()
		}
	}
	return ""
}
