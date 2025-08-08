package logger

import (
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	debugLogger *log.Logger
	errorLogger *log.Logger
	outputFile  *os.File
)

type ILogger interface {
	Info(message string)
	Debug(message string)
	Error(message string)
	Init(filePath string) error
	Close() error
}

// func NewLogger() ILogger {
// 	return &Logger{
// 		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
// 		debugLogger: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
// 		errorLogger: log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
// 	}
// }

// Init initializes the logger with an optional output file
// If filePath is provided and valid, all log output will be redirected to that file
// If filePath is empty or invalid, output will continue to stdout
func Init(filePath string) error {
	if filePath == "" {
		// Reset to stdout if no file path provided
		outputFile = nil
		infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
		errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		return nil
	}

	// Open file for writing (create if doesn't exist, append if exists)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	outputFile = file
	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

// Close closes the output file if one was opened
func Close() error {
	if outputFile != nil {
		return outputFile.Close()
	}
	return nil
}

func Info(message string) {
	infoLogger.Output(2, message)
}

func Debug(message string) {
	debugLogger.Output(2, message)
}

func Error(message string) {
	errorLogger.Output(2, message)
}
