package main

import (
	"gotest/mylogger"
)

var logger mylogger.Logger

func main() {
	f1()
}

func f1() {
	logger = mylogger.NewFileLogger("info", "./", "a.log")
	// logger = mylogger.NewConsoleLogger("info")
	defer logger.Close()
	logger.Debug("this debug")
	logger.Info("this info")
	logger.Fatal("this fatal")
}
