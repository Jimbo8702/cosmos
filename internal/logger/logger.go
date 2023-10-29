package logger

import (
	"github.com/sirupsen/logrus"
)

// LogLevel represents different log levels.
type LogLevel int

const (
    // ERROR level for error messages.
    ERROR LogLevel = iota
    // INFO level for informational messages.
    INFO
    // WARN level for warning messages.
    WARN
    // FATAL level for critical errors.
    FATAL
)

// Logger is an interface for logging messages at different log levels.
type Logger interface {
    Log(LogLevel, string, any)
}

type LogrusLogger struct {
	log *logrus.Logger
}

func NewLogrus() Logger {
	return &LogrusLogger{
		log: logrus.New(),
	}
}

func (ll *LogrusLogger) Log(level LogLevel, message string, data any) {
    switch level {
    case ERROR:
        ll.log.Error("MESSAGE=", message, " ERROR=", data)
    case INFO:
		ll.log.Info("MESSAGE=",message, " INFO=", data)
    case WARN:
        ll.log.Warn(message, data)
    case FATAL:
       ll.log.Fatal(message, data)
    }
}

