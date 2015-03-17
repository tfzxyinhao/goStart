package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

const (
	INFO = iota
	DEBUG
	WARN
	ERROR
)

type Logger struct {
	limit   int
	path    string
	console bool
	buffer  [4]bytes.Buffer
	prix    []string
}

func NewLogger(size int, path string) *Logger {
	c := make(chan os.Signal, 1)
	logger := &Logger{limit: size, path: path, console: true, prix: []string{"Info", "Debug", "Warn", "Error"}}
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		for sig := range c {
			fmt.Println(sig.String())
			signal.Stop(c)
			logger.shutdown()
			break
		}
	}()
	return logger
}

func (logger *Logger) savefile(index uint) {
	now := time.Now()
	y, m, d := now.Date()
	filename := fmt.Sprintf("%s_%02d%02d%02d_%02d%02d.log", logger.prix[index], y, m, d, now.Hour(), now.Second())
	file := filepath.Join(logger.path, filename)
	ioutil.WriteFile(file, logger.buffer[index].Bytes(), os.ModePerm)
	logger.buffer[index].Reset()
	fmt.Println(filename, file, logger.buffer[index].Len())
}

func (logger *Logger) save(index uint) {
	if index < 4 && logger.buffer[index].Len() >= logger.limit {
		logger.savefile(index)
	}
}

func (logger *Logger) write(index uint, args ...interface{}) {
	logger.buffer[index].WriteString(time.Now().Format(time.RFC3339))
	logger.buffer[index].WriteByte(' ')
	logger.buffer[index].WriteString(fmt.Sprint(args...))
	logger.buffer[index].WriteString("\n")
	if logger.console {
		fmt.Println(args...)
	}
	logger.save(index)
}

func (logger *Logger) shutdown() {
	var index uint = 0
	for index = 0; index < 4; index++ {
		logger.savefile(index)
	}
}

func (logger *Logger) Info(args ...interface{}) {
	logger.write(INFO, args...)
}

func (logger *Logger) Warn(args ...interface{}) {
	logger.write(WARN, args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.write(DEBUG, args...)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.write(ERROR, args...)
}
