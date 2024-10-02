package utils

import (
	"log"
	"os"
)

type Logger struct {
	file *os.File
}

// InitializeLogger initializes the logger to log to both a file and stdout.
func InitializeLogger(logFile string) (*Logger, error) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	log.SetOutput(file)
	return &Logger{file: file}, nil
}

// Info logs informational messages.
func (l *Logger) Info(msg string, args ...interface{}) {
	log.Printf("INFO: "+msg, args...)
}

// Warn logs warnings that require attention but are not errors.
func (l *Logger) Warn(msg string, args ...interface{}) {
	log.Printf("WARN: "+msg, args...)
}

// Error logs critical errors.
func (l *Logger) Error(msg string, args ...interface{}) {
	log.Printf("ERROR: "+msg, args...)
}

// Close closes the log file.
func (l *Logger) Close() {
	if l.file != nil {
		l.file.Close()
	}
}
