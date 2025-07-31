package logger

import (
	"log"
	"os"
)

type Logger struct {
	infoLogger  *log.Logger
	debugLogger *log.Logger
	errorLogger *log.Logger
	outputFile  *os.File
}

type ILogger interface {
	Info(message string)
	Debug(message string)
	Error(message string)
	Init(filePath string) error
	Close() error
}

func NewLogger() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLogger: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Init initializes the logger with an optional output file
// If filePath is provided and valid, all log output will be redirected to that file
// If filePath is empty or invalid, output will continue to stdout
func (l *Logger) Init(filePath string) error {
	if filePath == "" {
		// Reset to stdout if no file path provided
		l.outputFile = nil
		l.infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		l.debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
		l.errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		return nil
	}

	// Open file for writing (create if doesn't exist, append if exists)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	l.outputFile = file
	l.infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.debugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

// Close closes the output file if one was opened
func (l *Logger) Close() error {
	if l.outputFile != nil {
		return l.outputFile.Close()
	}
	return nil
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
