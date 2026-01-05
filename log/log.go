package log

import (
    "log"
    "os"
    "path/filepath"
)

type loggers struct {
    logInfo *log.Logger
    logWarn *log.Logger
    logErr  *log.Logger
}

var l loggers

func init() {
    
    logDir := filepath.Join(os.Getenv("HOME"), "Projects", "go_project", "logfile")

    
    if err := os.MkdirAll(logDir, 0755); err != nil {
        log.Fatalf("cannot create log dir: %v", err)
    }

    flags := log.LstdFlags | log.Lshortfile

    fileInfo, _ := os.OpenFile(filepath.Join(logDir, "log_info.log"),
        os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    fileWarn, _ := os.OpenFile(filepath.Join(logDir, "log_warn.log"),
        os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    fileErr, _ := os.OpenFile(filepath.Join(logDir, "log_err.log"),
        os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

    l = loggers{
        logInfo: log.New(fileInfo, "INFO:\t", flags),
        logWarn: log.New(fileWarn, "WARN:\t", flags),
        logErr:  log.New(fileErr, "ERR:\t", flags),
    }
}

func Info(v ...interface{}) { l.logInfo.Println(v...) }
func Warn(v ...interface{}) { l.logWarn.Println(v...) }
func Err(v ...interface{})  { l.logErr.Println(v...) }
