package rlog

import (
	"context"
	"os"
	"testing"

	"github.com/goqianjin/common-libs/xlog/internal"
)

func TestRLog(t *testing.T) {
	log := New(os.Stdout, Option{})
	log.Log(context.Background(), internal.LevelTrace, "raw log trace info", "arg1", "arg2")
	log.Log(context.Background(), internal.LevelDebug, "raw log debug info", "arg1", "arg2")
	log.Log(context.Background(), internal.LevelInfo, "raw log info info", "arg1", "arg2")
	log.Log(context.Background(), internal.LevelWarn, "raw log warn info", "arg1", "arg2")
	log.Log(context.Background(), internal.LevelError, "raw log error info", "arg1", "arg2")
	log.Log(context.Background(), internal.LevelFatal, "raw log fatal info", "arg1", "arg2")
}
