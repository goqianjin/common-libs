package xlog

import (
	"context"
	"os"
	"testing"

	"github.com/goqianjin/common-libs/xlog/internal"
	"github.com/goqianjin/common-libs/xlog/slog"
)

func TestPackageLog(t *testing.T) {
	//SetLevel(LevelTrace)
	// default format: classical
	Trace(context.Background(), "Trace information in classical format.", "key1", "value1")
	Debug(context.Background(), "Debug information in classical format.", "key1", "value1")
	Info(context.Background(), "Info information in classical format.", "key1", "value1")
	Warn(context.Background(), "Warn information in classical format.", "key1", "value1")
	Error(context.Background(), "Error information in classical format.", "key1", "value1")
	//Fatal(context.Background(), "Fatal information in classical format.", "key1", "value1")

	// format: json
	SetDefault(slog.New(internal.DefaultW, slog.Option{Format: FormatJSON}))
	Trace(context.Background(), "Trace information in JSON format.", "key1", "value1")
	Debug(context.Background(), "Debug information in JSON format.", "key1", "value1")
	Info(context.Background(), "Info information in JSON format.", "key1", "value1")
	Warn(context.Background(), "Warn information in JSON format.", "key1", "value1")
	Error(context.Background(), "Error information in JSON format.", "key1", "value1")
	//Fatal(context.Background(), "Fatal information in JSON format.", "key1", "value1")

	// format: text
	SetDefault(slog.New(internal.DefaultW, slog.Option{Format: FormatText}))
	Trace(context.Background(), "Trace information in text format.", "key1", "value1")
	Debug(context.Background(), "Debug information in text format.", "key1", "value1")
	Info(context.Background(), "Info information in text format.", "key1", "value1")
	Warn(context.Background(), "Warn information in text format.", "key1", "value1")
	Error(context.Background(), "Error information in text format.", "key1", "value1")
	//Fatal(context.Background(), "Fatal information in text format.", "key1", "value1")

	// log with xlog context
	_, ctx := NewSLog(os.Stdout, SlogOption{ReqID: "contextLogID", Format: FormatClassical})
	Trace(ctx, "Trace package context information in classical format.", "key1", "value1")
	Debug(ctx, "Debug package context information in classical format.", "key1", "value1")
	Info(ctx, "Info package context information in classical format.", "key1", "value1")
	Warn(ctx, "Warn package context information in classical format.", "key1", "value1")
	Error(ctx, "Error package context information in classical format.", "key1", "value1")
	//Fatal(ctx, "Fatal package context information in classical format.", "key1", "value1")
}

func TestInstanceLog(t *testing.T) {
	log, _ := NewSLog(os.Stdout, SlogOption{})
	log.Trace("Trace instance context information in classical format.", "key1", "value1")
	log.Debug("Debug instance context information in classical format.", "key1", "value1")
	log.Info("Info instance context information in classical format.", "key1", "value1")
	log.Warn("Warn instance context information in classical format.", "key1", "value1")
	log.Error("Error instance context information in classical format.", "key1", "value1")
	//log.Fatal("Fatal instance context information in classical format.", "key1", "value1")
}

func TestPackageLogf(t *testing.T) {
	level := LevelTrace
	SetDefault(slog.New(internal.DefaultW, slog.Option{Level: &level}))
	// default format: classical
	Tracef(context.Background(), "Trace information in classical format. {%s: %v}", "key1", "value1")
	Debugf(context.Background(), "Debug information in classical format. {%s: %v}", "key1", "value1")
	Infof(context.Background(), "Info information in classical format. {%s: %v}", "key1", "value1")
	Warnf(context.Background(), "Warn information in classical format. {%s: %v}", "key1", "value1")
	Errorf(context.Background(), "Error information in classical format. {%s: %v}", "key1", "value1")
	//Fatalf(context.Background(), "Fatal information in classical format. {%s: %v}", "key1", "value1")

	// format: json
	SetDefault(slog.New(internal.DefaultW, slog.Option{Format: FormatJSON}))
	Tracef(context.Background(), "Trace information in JSON format. {%s: %v}", "key1", "value1")
	Debugf(context.Background(), "Debug information in JSON format. {%s: %v}", "key1", "value1")
	Infof(context.Background(), "Info information in JSON format. {%s: %v}", "key1", "value1")
	Warnf(context.Background(), "Warn information in JSON format. {%s: %v}", "key1", "value1")
	Errorf(context.Background(), "Error information in JSON format. {%s: %v}", "key1", "value1")
	//Fatalf(context.Background(), "Fatal information in JSON format. {%s: %v}", "key1", "value1")

	// format: text
	SetDefault(slog.New(internal.DefaultW, slog.Option{Format: FormatText}))
	Tracef(context.Background(), "Trace information in text format. {%s: %v}", "key1", "value1")
	Debugf(context.Background(), "Debug information in text format. {%s: %v}", "key1", "value1")
	Infof(context.Background(), "Info information in text format. {%s: %v}", "key1", "value1")
	Warnf(context.Background(), "Warn information in text format. {%s: %v}", "key1", "value1")
	Errorf(context.Background(), "Error information in text format. {%s: %v}", "key1", "value1")
	//Fatalf(context.Background(), "Fatal information in text format. {%s: %v}", "key1", "value1")

	// log with xlog context
	_, ctx := NewSLog(os.Stdout, SlogOption{ReqID: "contextLogID", Format: FormatClassical})
	Tracef(ctx, "Trace package context information in classical format. {%s: %v}", "key1", "value1")
	Debugf(ctx, "Debug package context information in classical format. {%s: %v}", "key1", "value1")
	Infof(ctx, "Info package context information in classical format. {%s: %v}", "key1", "value1")
	Warnf(ctx, "Warn package context information in classical format. {%s: %v}", "key1", "value1")
	Errorf(ctx, "Error package context information in classical format. {%s: %v}", "key1", "value1")
	//Fatalf(ctx, "Fatal package context information in classical format. {%s: %v}", "key1", "value1")
}

func TestInstanceLogf(t *testing.T) {
	log, _ := NewSLogf(os.Stdout, SlogOption{})
	log.Tracef("Trace instance context information in classical format. {%s: %v}", "key1", "value1")
	log.Debugf("Debug instance context information in classical format. {%s: %v}", "key1", "value1")
	log.Infof("Info instance context information in classical format. {%s: %v}", "key1", "value1")
	log.Warnf("Warn instance context information in classical format. {%s: %v}", "key1", "value1")
	log.Errorf("Error instance context information in classical format. {%s: %v}", "key1", "value1")
	//log.Fatalf("Fatal instance context information in classical format. {%s: %v}", "key1", "value1")
}
