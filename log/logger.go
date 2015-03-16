package logger

import (
	"bytes"
)

type Logger struct {
	limit  uint
	path   string
	buffer bytes.Buffer
}

func NewLogger(size uint, path string) *Logger {
	return &Logger{}
}

func (logger *Logger) Info() {

}

func (logger *Logger) Warn() {

}

func (logger *Logger) Debug() {

}

func (logger *Logger) Error() {

}
