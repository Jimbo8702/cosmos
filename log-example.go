package main

import (
	"Jimbo8702/randomThoughts/cosmos/cosmos"
	"Jimbo8702/randomThoughts/cosmos/internal/logger"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type MyLogger struct {
	level logger.LogLevel
}

func (s *MyLogger) Log(message string, data any) {
	var strLevel string
	lvl, err := getLevelStr(s.level)
	if err != nil {
		strLevel = "INFO"
	} else {
		strLevel = lvl
	}
	fmt.Printf("%s: message=%s, data=%v\n", strLevel, message, data)
}

func (s *MyLogger) LogWithLevel(level logger.LogLevel, message string, data any) {
	var strLevel string
	lvl, err := getLevelStr(level)
	if err != nil {
		strLevel = "INFO"
	} else {
		strLevel = lvl
	}
	fmt.Printf("%s: message=%s, data=%v\n", strLevel, message, data)
}

func (s *MyLogger) GetLogWriter() io.Writer {
	return os.Stdout
}

func getLevelStr(level logger.LogLevel) (string, error) {
	switch level {
    case logger.ERROR:
        return "ERROR", nil
    case logger.INFO:
		return "INFO", nil
    case logger.WARN:
		return "WARN", nil
    case logger.FATAL:
        return "FATAL", nil
    default:
        return "", errors.New("level not supported")
    }
}

func LogExample() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	app, err := cosmos.New(path)
	if err != nil {
		log.Fatal(err)
	}
	
	app.ErrorLog.Log("this is an example error with default(logrus)", nil)
	app.InfoLog.Log("this is an example Info with default(logrus)", nil)

	// an example on how you can use different loggers 
	standardErrorLog := &MyLogger{
		level: logger.ERROR,
	}
	app.ErrorLog = standardErrorLog
	app.ErrorLog.Log("this is now using the format package from the standard libaray", nil)

	standardInfoLog := &MyLogger{
		level: logger.INFO,
	}
	app.InfoLog = standardInfoLog
	app.InfoLog.Log("this is now using the format package from standard libaray", nil)
}