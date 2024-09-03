package rlog

import (
	"io"
	"sync"

	"github.com/goqianjin/common-libs/xlog/internal"
)

func New(w io.Writer, option Option) internal.LoggerInternal {
	mutex := &sync.Mutex{}
	if w == internal.DefaultW {
		mutex = &internal.DefaultWMutex
	}

	if option.Level == nil {
		option.Level = &internal.DefaultLevel
	}
	if option.Separator == "" {
		option.Separator = "\t"
	}

	return &rawLogger{
		w:         w,
		mu:        mutex,
		level:     *option.Level,
		separator: option.Separator,
	}
}

// ------ options ------

type Option struct {
	Level *internal.Level // Optional: default 0 is INFO

	// TODO: support Format such as JSON

	Separator string // Optional: default is '	' (tab character)

	// Optional: field names. default is each context attributes, msg, args...
	// (1) 'msg' and 'args' are automatically appended in the fields names if unspecified
	// (2) these values of custom field should be put into context attributes
	// before call Log method, the default is '-' if any value is missing.
	FieldNames []string
}
