package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

const MAX_LOG_FILE_SIZE = 1024 * 1024 * 10 //单文件最大10M

const (
	DEBUG = iota
	NORMAL
	ERROR
	FATAL
)

//rotatable log
var lm LogManager

type LogManager struct {
	logger   *log.Logger
	file     *os.File
	fileName string
}

func Init(fileName string) bool {
	lm.fileName = fileName
	lm.file, lm.logger = newLogger()
	return lm.file == nil || lm.logger == nil
}

func Info(str string) {
	lm.logger.SetPrefix("[NORMAL]")
	Write(str)
}

func Debug(str string) {
	lm.logger.SetPrefix("[DEBUG]")
	Write(str)
}

func Error(str string) {
	lm.logger.SetPrefix("[ERROR]")
	Write(str)
}

func Fatal(str string) {
	lm.logger.SetPrefix("[FATAL]")
	Write(str)
}

func Write(str string) {

	lm.logger.Println(str)

	fi, err := lm.file.Stat()
	if err != nil {
		fmt.Println("get log file state failed")
	} else {
		//rotate
		if fi.Size() > MAX_LOG_FILE_SIZE {
			lm.file.Close()
			file, logger := newLogger()
			if file != nil && logger != nil {
				lm.file = file
				lm.logger = logger
			}
		}
	}
}

func newLogger() (*os.File, *log.Logger) {
	var logFile *os.File
	var logger *log.Logger
	fileName := fmt.Sprintf("%s_%s.log", lm.fileName, time.Now().Format("2006-01-02_15-04-05"))
	logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("create log file failed")
	} else {
		logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
	}
	return logFile, logger
}
