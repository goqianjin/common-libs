package internal

import (
	"context"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

func SetGenReqIDFunc(f func() string) {
	if f == nil {
		f = defaultGenReqId
	}
	genReqID = f
}

func GenReqID() string {
	return genReqID()
}

// ------ helpers ------
// ReqID implementation
// @see https://github.com/qbox/kodo/blob/develop/libs/xlog.v1/xlog.go#L74-L97

var pid = uint32(os.Getpid())

var genReqID = defaultGenReqId

func defaultGenReqId() string {
	var b [12]byte
	binary.LittleEndian.PutUint32(b[:], pid)
	binary.LittleEndian.PutUint64(b[4:], uint64(time.Now().UnixNano()))
	return base64.URLEncoding.EncodeToString(b[:])
}

// ------ utils ------

var AttributeKeyReqID = "__reqID__"

func GetContextReqID(ctx context.Context) string {
	v := GetContextAttribute(ctx, AttributeKeyReqID)
	switch v.(type) {
	case string:
		return v.(string)
	default:
		return fmt.Sprint(v)
	}
}
