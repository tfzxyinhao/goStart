package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLogger(t *testing.T) {
	curDir, err := os.Getwd()
	dir := filepath.Join(curDir, "log")
	if err != nil {
		t.Error(err)
	} else {
		logger := NewLogger(100, dir)
		for i := 0; i < 10; i++ {
			logger.Info("test1")
		}
	}
}
