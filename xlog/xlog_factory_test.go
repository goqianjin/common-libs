package xlog

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/goqianjin/common-libs/xlog/internal"
)

func TestWrapLogger(t *testing.T) {
	logger, ctx := WrapLogger(context.Background(), &goLoggerMock{
		logger: log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile),
	})

	// instance logger
	logger.Trace("Trace instance information in wrapping go log, %s=%v", "key1", "value1")
	logger.Debug("Debug instance information in wrapping go log, %s=%v", "key1", "value1")
	logger.Info("Info instance information in wrapping go log, %s=%v", "key1", "value1")
	logger.Warn("Warn instance information in wrapping go log, %s=%v", "key1", "value1")
	logger.Error("Error instance information in wrapping go log, %s=%v", "key1", "value1")
	//logger.Fatal("Fatal instance information in wrapping go log, %s=%v", "key1", "value1")

	// package logger
	Trace(ctx, "Trace package information in wrapping go log, %s=%v", "key1", "value1")
	Debug(ctx, "Debug package information in wrapping go log, %s=%v", "key1", "value1")
	Info(ctx, "Info package information in wrapping go log, %s=%v", "key1", "value1")
	Warn(ctx, "Warn package information in wrapping go log, %s=%v", "key1", "value1")
	Error(ctx, "Error package information in wrapping go log, %s=%v", "key1", "value1")
	//Fatal(ctx, "Fatal package information in wrapping go log, %s=%v", "key1", "value1")
}

func TestWrapLogf(t *testing.T) {
	logger, ctx := WrapLogf(context.Background(), &goLoggerMock{
		logger: log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile),
	})

	// instance logger
	logger.Tracef("Tracef instance information in wrapping go log, %s=%v", "key1", "value1")
	logger.Debugf("Debugf instance information in wrapping go log, %s=%v", "key1", "value1")
	logger.Infof("Infof instance information in wrapping go log, %s=%v", "key1", "value1")
	logger.Warnf("Warnf instance information in wrapping go log, %s=%v", "key1", "value1")
	logger.Errorf("Errorf instance information in wrapping go log, %s=%v", "key1", "value1")
	//logger.Fatalf("Fatalf instance information in wrapping go log, %s=%v", "key1", "value1")

	// package logger
	Tracef(ctx, "Tracef package information in wrapping go log, %s=%v", "key1", "value1")
	Debugf(ctx, "Debugf package information in wrapping go log, %s=%v", "key1", "value1")
	Infof(ctx, "Infof package information in wrapping go log, %s=%v", "key1", "value1")
	Warnf(ctx, "Warnf package information in wrapping go log, %s=%v", "key1", "value1")
	Errorf(ctx, "Errorf package information in wrapping go log, %s=%v", "key1", "value1")
	//Fatalf(ctx, "Fatal package information in wrapping go log, %s=%v", "key1", "value1")
}

// ------ goLoggerMock ------

type goLoggerMock struct {
	logger *log.Logger
}

func (l *goLoggerMock) Log(ctx context.Context, level internal.Level, msg string, args ...any) {
	// call depth: [log.(*Logger).Output, xlog/slog.(*loggerAdapter).Log,
	// this function, loggerAdapter.<Function> or xlog package caller]
	l.logger.Output(4, fmt.Sprintf(msg, args...))
	switch level {
	case LevelFatal:
		os.Exit(1)
	}
}
