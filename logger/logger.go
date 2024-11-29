package logger

import (
	"fmt"
	"time"
)

type LoggerIF interface {
	INFO(tag string, message string)
	ERROR(tag string, message string)
}

type Logger struct {
}

func NewLogger() *Logger {

	return &Logger{}
}

func (s *Logger) print(level string, tag string, message string) {

	currentTime := time.Now().Format("2006-01-02 15:04:05.000")
	logMessage := fmt.Sprintf("[%s][%s][%s]: %s", currentTime, level, tag, message)
	fmt.Println(logMessage)
}
func (s *Logger) INFO(tag string, message string) {
	s.print("INFO", tag, message)
}

func (s *Logger) ERROR(tag string, message string) {
	s.print("ERROR", tag, message)
}
