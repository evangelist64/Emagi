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
	msgChan  chan string
	fileName string
	isDebug  bool
}

func Init(fileName string, isDebug bool) bool {
	lm.fileName = fileName
	lm.isDebug = isDebug
	lm.msgChan = make(chan string, 100)
	lm.file, lm.logger = newLogger()
	return lm.file == nil || lm.logger == nil
}

func Run() {

	for {
		str := <-lm.msgChan
		//控制台输出
		if lm.isDebug {
			log.Println(str)
		}
		//文件输出
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
}

func Info(str string) {
	Write("[NORMAL]" + str)
}

func Debug(str string) {
	if lm.isDebug {
		Write("[DEBUG]" + str)
	}
}

func Error(str string) {
	Write("[ERROR]" + str)
}

func Fatal(str string) {
	Write("[FATAL]" + str)
}

func Write(str string) {
	lm.msgChan <- str
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
