package logger

import "fmt"


type StdLogger struct {
	level LogLevel
}

func (s *StdLogger) SetLevel(level LogLevel) {
	s.level = level
}

func (s *StdLogger) Log(message string, data any) {
	var strLevel string
	lvl, err := getLevelStr(s.level)
	if err != nil {
		strLevel = "INFO"
	} else {
		strLevel = lvl
	}
	fmt.Printf("%s: message=%s, data=%v\n", strLevel, message, data)
}

func (s *StdLogger) LogWithLevel(level LogLevel, message string, data any) {
	var strLevel string
	lvl, err := getLevelStr(level)
	if err != nil {
		strLevel = "INFO"
	} else {
		strLevel = lvl
	}
	fmt.Printf("%s: message=%s, data=%v\n", strLevel, message, data)
}
