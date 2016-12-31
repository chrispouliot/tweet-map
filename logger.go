package main

import (
	"log"
	"os"
)

type Logger struct {
	info *log.Logger
	err  *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		info: log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime),
		err:  log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime),
	}
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
