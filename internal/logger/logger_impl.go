package logger

import "fmt"

type Log struct {
	log 	Logger
	level 	LogLevel
}
//Default logger set to logrus
func New(level LogLevel) Logger {
	log := NewLogrus()
	return &Log{
		log: log,
		level: level,
	}
}

func (l *Log) Log(message string, data any) {
	l.log.LogWithLevel(l.level, message, data)
}

func (l *Log) LogWithLevel(level LogLevel, message string, data any) {
	l.log.LogWithLevel(level, message, data)
}

type StdLogger struct {}

func (s *StdLogger) Log(message string, data any) {
	fmt.Printf("message=%s, data=%v", message, data)
}

func (s *StdLogger) LogWithLevel(level LogLevel, message string, data any) {
	var strLevel string
	switch level {
    case ERROR:
        strLevel = "ERROR"
    case INFO:
		strLevel = "INFO"
    case WARN:
		strLevel = "WARN"
    case FATAL:
        strLevel = "FATAL"
    }
	fmt.Printf("%s: message=%s, data=%v", strLevel, message, data)
}
