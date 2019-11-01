package log

import (
    "os"
    "testing"
)

var logger = NewLogger(os.Stdout)

func TestDebug(t *testing.T) {
    logger.SetLevel("debug")
    logger.Debug("debug")

    SetLevel("debug")
    Debug("debug")
}

func TestDebugf(t *testing.T) {
    logger.SetLevel("debug")
    logger.Debugf("debugf")

    SetLevel("debug")
    Debugf("debugf")
}

func TestInfo(t *testing.T) {
    logger.SetLevel("info")
    logger.Info("info")

    SetLevel("info")
    Info("info")
}

func TestInfof(t *testing.T) {
    logger.SetLevel("info")
    logger.Infof("infof")

    SetLevel("info")
    Infof("infof")
}

func TestWarn(t *testing.T) {
    logger.SetLevel("warn")
    logger.Warn("warn")

    SetLevel("warn")
    Warn("warn")
}

func TestWarnf(t *testing.T) {
    logger.SetLevel("warn")
    logger.Warnf("warnf")

    SetLevel("warn")
    Warnf("warnf")
}

func TestError(t *testing.T) {
    logger.SetLevel("error")
    logger.Error("error")

    SetLevel("error")
    Error("error")
}

func TestErrorf(t *testing.T) {
    logger.SetLevel("error")
    logger.Errorf("errorf")

    SetLevel("error")
    Errorf("errorf")
}

func TestGetLevel(t *testing.T) {
    if getLevel("debug") != DEBUG {
        t.FailNow()

        return
    }

    if getLevel("info") != INFO {
        t.FailNow()

        return
    }

    if getLevel("warn") != WARN {
        t.FailNow()

        return
    }

    if getLevel("error") != ERROR {
        t.FailNow()

        return
    }
}

func TestLoggerSetLevel(t *testing.T) {
    logger.SetLevel("debug")

    if logger.level != DEBUG {
        t.FailNow()

        return
    }
}

func TestIsDebugEnabled(t *testing.T) {
    logger.SetLevel("debug")

    if !logger.IsDebugEnabled() {
        t.FailNow()

        return
    }
}

func TestIsWarnEnabled(t *testing.T) {
    logger.SetLevel("warn")

    if !logger.IsWarnEnabled() {
        t.FailNow()

        return
    }
}
