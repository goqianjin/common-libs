
#### log 静态内容

```go
func New(out io.Writer, prefix string, flag int) *Logger {
    l := new(Logger)
    l.SetOutput(out)
    l.SetPrefix(prefix)
    l.SetFlags(flag)
    return l
}

func Default() *Logger { return std }

var std = New(os.Stderr, "", LstdFlags)

func Flags() int {
    return std.Flags()
}

func SetFlags(flag int) {
    std.SetFlags(flag)
}

func Prefix() string {
    return std.Prefix()
}

func SetPrefix(prefix string) {
    std.SetPrefix(prefix)
}

func Writer() io.Writer {
    return std.Writer()
}


func SetOutput(w io.Writer) {
    std.SetOutput(w)
}

func Output(calldepth int, s string) error {
    return std.Output(calldepth+1, s) // +1 for this frame.
}
func Print(v ...any) {
    std.output(0, 2, func(b []byte) []byte {
        return fmt.Append(b, v...)
    })
}
```

#### slog 静态内容

```go
var defaultLogger atomic.Pointer[Logger]

func Default() *Logger { return defaultLogger.Load() }

func SetDefault(l *Logger) {
    defaultLogger.Store(l)
    // If the default's handler is a defaultHandler, then don't use a handleWriter,
    // or we'll deadlock as they both try to acquire the log default mutex.
    // The defaultHandler will use whatever the log default writer is currently
    // set to, which is correct.
    // This can occur with SetDefault(Default()).
    // See TestSetDefault.
    if _, ok := l.Handler().(*defaultHandler); !ok {
        capturePC := log.Flags()&(log.Lshortfile|log.Llongfile) != 0
        log.SetOutput(&handlerWriter{l.Handler(), &logLoggerLevel, capturePC})
        log.SetFlags(0) // we want just the log message, no time or location
    }
}

func New(h Handler) *Logger {
    if h == nil {
        panic("nil Handler")
    }
    return &Logger{handler: h}
}

func With(args ...any) *Logger {
    return Default().With(args...)
}

func NewLogLogger(h Handler, level Level) *log.Logger {
    return log.New(&handlerWriter{h, level, true}, "", 0)
}


func NewJSONHandler(w io.Writer, opts *HandlerOptions) *JSONHandler {
    if opts == nil {
        opts = &HandlerOptions{}
    }
    return &JSONHandler{
        &commonHandler{
            json: true,
            w:    w,
            opts: *opts,
            mu:   &sync.Mutex{},
        },
    }
}



var logLoggerLevel LevelVar

func SetLogLoggerLevel(level Level) (oldLevel Level) {
    oldLevel = logLoggerLevel.Level()
    logLoggerLevel.Set(level)
    return
}

// Info calls [Logger.Info] on the default logger.
func Info(msg string, args ...any) {
    Default().log(context.Background(), LevelInfo, msg, args...)
}

```

