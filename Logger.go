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

func (logger *Logger) Save(index uint) {
	if index < 4 && logger.buffer[index].Len() >= logger.limit {
		now := time.Now()
		y, m, d := now.Date()
		filename := fmt.Sprintf("%s_%2d%2d%2d_%2d%2d.log", logger.prix[index], y, m, d, now.Hour(), now.Second())
		file := filepath.Join(logger.path, filename)
		ioutil.WriteFile(file, logger.buffer[index].Bytes(), os.ModePerm)
		logger.buffer[index].Reset()
		fmt.Println(filename, file, logger.buffer[index].Len())
	}
}

func (logger *Logger) shutdown() {
	for index := 0; index < 4; index++ {
		now := time.Now()
		y, m, d := now.Date()
		filename := fmt.Sprintf("%s_%2d%2d%2d_%2d%2d.log", logger.prix[index], y, m, d, now.Hour(), now.Second())
		file := filepath.Join(logger.path, filename)
		ioutil.WriteFile(file, logger.buffer[index].Bytes(), os.ModePerm)
		logger.buffer[index].Reset()
		fmt.Println(filename, file, logger.buffer[index].Len())
	}
}

func (logger *Logger) Info(args ...interface{}) {
	logger.buffer[INFO].WriteString(fmt.Sprint(args))
	if logger.console {
		fmt.Println(args)
	}
	logger.Save(INFO)
}

func (logger *Logger) Warn(args ...interface{}) {
	logger.buffer[WARN].WriteString(fmt.Sprint(args))
	if logger.console {
		fmt.Println(args)
	}
	logger.Save(INFO)
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.buffer[DEBUG].WriteString(fmt.Sprint(args))
	if logger.console {
		fmt.Println(args)
	}
	logger.Save(INFO)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.buffer[ERROR].WriteString(fmt.Sprint(args))
	if logger.console {
		fmt.Println(args)
	}
	logger.Save(INFO)
}
