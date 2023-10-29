package logger

import "io"

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

func (l *Log) GetLogWriter() io.Writer {
	return l.log.GetLogWriter()
}
