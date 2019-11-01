package log

import (
    "fmt"
    "io"
    stdlog "log"
    "os"
    "strings"
)

const (
    _ = iota
    DEBUG
    INFO
    WARN
    ERROR
    FATAL
)

var logLevel = DEBUG

type Logger struct {
    level  int
    logger *stdlog.Logger
}

var std = NewLogger(os.Stdout)

func NewLogger(out io.Writer) *Logger {
    return &Logger{level: logLevel, logger: stdlog.New(out, "", stdlog.Ldate|stdlog.Ltime|stdlog.Lshortfile)}
}

func getLevel(level string) int {
    level = strings.ToLower(level)

    switch level {
    case "debug":
        return DEBUG
    case "info":
        return INFO
    case "warn":
        return WARN
    case "error":
        return ERROR
    case "fatal":
        return FATAL
    default:
        return INFO
    }
}

func (l *Logger) SetLevel(level string) {
    l.level = getLevel(level)
}

func (l *Logger) IsDebugEnabled() bool {
    return l.level <= DEBUG
}

func (l *Logger) IsWarnEnabled() bool {
    return l.level <= WARN
}

func (l *Logger) Debug(v ...interface{}) {
    if DEBUG < l.level {
        return
    }

    l.logger.SetPrefix("[DEBUG] ")
    _ = l.logger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
    if DEBUG < l.level {
        return
    }

    l.logger.SetPrefix("[DEBUG] ")
    _ = l.logger.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
    if INFO < l.level {
        return
    }

    l.logger.SetPrefix("[INFO] ")
    _ = l.logger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
    if INFO < l.level {
        return
    }

    l.logger.SetPrefix("[INFO] ")
    _ = l.logger.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
    if WARN < l.level {
        return
    }

    l.logger.SetPrefix("[WARN] ")
    _ = l.logger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
    if WARN < l.level {
        return
    }

    l.logger.SetPrefix("[WARN] ")
    _ = l.logger.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
    if ERROR < l.level {
        return
    }

    l.logger.SetPrefix("[ERROR] ")
    _ = l.logger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
    if ERROR < l.level {
        return
    }

    l.logger.SetPrefix("[ERROR] ")
    _ = l.logger.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
    if FATAL < l.level {
        return
    }

    l.logger.SetPrefix("[FATAL] ")
    _ = l.logger.Output(2, fmt.Sprint(v...))
    os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
    if FATAL < l.level {
        return
    }

    l.logger.SetPrefix("[FATAL] ")
    _ = l.logger.Output(2, fmt.Sprintf(format, v...))
    os.Exit(1)
}

func SetLevel(level string) {
    std.level = getLevel(level)
}

func Debug(v ...interface{}) {
    if DEBUG < std.level {
        return
    }

    std.logger.SetPrefix("[DEBUG] ")
    _ = std.logger.Output(2, fmt.Sprint(v...))
}

func Debugf(format string, v ...interface{}) {
    if DEBUG < std.level {
        return
    }

    std.logger.SetPrefix("[DEBUG] ")
    _ = std.logger.Output(2, fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
    if INFO < std.level {
        return
    }

    std.logger.SetPrefix("[INFO] ")
    _ = std.logger.Output(2, fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
    if INFO < std.level {
        return
    }

    std.logger.SetPrefix("[INFO] ")
    _ = std.logger.Output(2, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
    if WARN < std.level {
        return
    }

    std.logger.SetPrefix("[WARN] ")
    _ = std.logger.Output(2, fmt.Sprint(v...))
}

func Warnf(format string, v ...interface{}) {
    if WARN < std.level {
        return
    }

    std.logger.SetPrefix("[WARN] ")
    _ = std.logger.Output(2, fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
    if ERROR < std.level {
        return
    }

    std.logger.SetPrefix("[ERROR] ")
    _ = std.logger.Output(2, fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
    if ERROR < std.level {
        return
    }

    std.logger.SetPrefix("[ERROR] ")
    _ = std.logger.Output(2, fmt.Sprintf(format, v...))
}

func Fatal(v ...interface{}) {
    if FATAL < std.level {
        return
    }

    std.logger.SetPrefix("[FATAL] ")
    _ = std.logger.Output(2, fmt.Sprint(v...))
    os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
    if FATAL < std.level {
        return
    }

    std.logger.SetPrefix("[FATAL] ")
    _ = std.logger.Output(2, fmt.Sprintf(format, v...))
    os.Exit(1)
}
