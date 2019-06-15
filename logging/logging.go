package logging

import (
	configs "catdogs-service/configs/common"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

const (
	INFO  string = "INFO"
	WARN  string = "WARN"
	ERROR string = "ERROR"
)

var (
	logger    *log.Logger
	logPrefix = ""
)

func init() {
	logFile, err := os.OpenFile(configs.C.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	logger = log.New(logFile, "", log.LstdFlags)
}

func setPrefix(level string) {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", level, filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", level)
	}
	logger.SetPrefix(logPrefix)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}

func Infof(format string, v ...interface{}) {
	setPrefix(INFO)
	logger.Printf(format, v...)
}

func Warn(v ...interface{}) {
	setPrefix(WARN)
	logger.Println(v...)
}

func Warnf(format string, v ...interface{}) {
	setPrefix(WARN)
	logger.Printf(format, v...)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}

func Errorf(format string, v ...interface{}) {
	setPrefix(ERROR)
	logger.Printf(format, v...)
}
