package logger

// Error logger
type ErrorLog struct {
	log    Logger
	level  LogLevel
}

func (el *ErrorLog) Log(message string, data any) {
	el.log.LogWithLevel(el.level, message, data)
}

func (el *ErrorLog) LogWithLevel(level LogLevel, message string, data any) {
	el.log.LogWithLevel(level, message, data)
}

func NewErrorLog() Logger {
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
	il.log.LogWithLevel(il.level, message, data)
}
func (il *InfoLog) LogWithLevel(level LogLevel, message string, data any) {
	il.log.LogWithLevel(level, message, data)
}

func NewInfoLog() Logger {
	log := NewLogrus()
	return &InfoLog{
		log: log,
		level: INFO,
	}
}




