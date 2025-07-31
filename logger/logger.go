package logger

import (
	"log"
	"os"
)

type Logger struct {
	infoLogger  *log.Logger
	debugLogger *log.Logger
	errorLogger *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLogger: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(message string) {
	l.infoLogger.Output(2, message)
}

func (l *Logger) Debug(message string) {
	l.debugLogger.Output(2, message)
}

func (l *Logger) Error(message string) {
	l.errorLogger.Output(2, message)
}
