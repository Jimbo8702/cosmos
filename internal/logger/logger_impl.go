package logger

// Error logger
type ErrorLog struct {
	log    Logger
	level  LogLevel
}

func (el *ErrorLog) Log(message string, data any) {
	el.log.Log(el.level, message, data)
}

func NewErrorLog() *ErrorLog {
	log := NewLogrus()
	return &ErrorLog{
		log: log,
		level: ERROR,
	}
}

// Info logger
type InfoLog struct {
	log    Logger
	level  LogLevel
}

func (il *InfoLog) Log(message string, data any) {
	il.log.Log(il.level, message, data)
}

func NewInfoLog() *InfoLog {
	log := NewLogrus()
	return &InfoLog{
		log: log,
		level: INFO,
	}
}
