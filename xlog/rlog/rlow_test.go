package rlog

import (
	"context"
	"os"
	"os/user"
	"testing"
	"time"

	"github.com/goqianjin/common-libs/xlog/internal"
)

func TestRLog(t *testing.T) {
	log := New(os.Stdout, Option{})

	// logging - basic
	log.Log(context.Background(), internal.LevelTrace, "basic RLog by TRACE info")
	log.Log(context.Background(), internal.LevelDebug, "basic RLog by DEBUG info")
	log.Log(context.Background(), internal.LevelInfo, "basic RLog by INFO info")
	log.Log(context.Background(), internal.LevelWarn, "basic RLog by WARN info")
	log.Log(context.Background(), internal.LevelError, "basic RLog by ERROR info")
	log.Log(context.Background(), internal.LevelFatal, "basic RLog by FATAL info")

	// logging with arguments
	log.Log(context.Background(), internal.LevelTrace, "RLog with args by TRACE info", "arg1", "arg2")
	log.Log(context.Background(), internal.LevelDebug, "RLog with args by DEBUG info", "arg1", "arg2")
	log.Log(context.Background(), internal.LevelInfo, "RLog with args by INFO info", "arg1", "arg2")
	log.Log(context.Background(), internal.LevelWarn, "RLog with args by WARN info", "arg1", "arg2")
	log.Log(context.Background(), internal.LevelError, "RLog with args by ERROR info", "arg1", "arg2")
	log.Log(context.Background(), internal.LevelFatal, "RLog with args by FATAL info", "arg1", "arg2")

	// logging - with attributes
	ctx := internal.PutContextAttribute(context.Background(), "timeAttr1", time.Now())
	ctx = internal.PutContextAttribute(ctx, "stringAttr1", "value of key1")
	ctx = internal.PutContextAttribute(ctx, "intAttr1", 123)
	ctx = internal.PutContextAttribute(ctx, "boolAttr1", true)
	ctx = internal.PutContextAttribute(ctx, "objectAtt1", user.User{Uid: "10001", Username: "xlog owner"})
	log.Log(ctx, internal.LevelInfo, "RLog with attributes and args by INFO info", "arg1", "arg2")
}
