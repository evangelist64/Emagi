package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

const MAX_LOG_FILE_SIZE = 1024 * 1024 * 10 //单文件最大10M

//rotatable log
var lm LogManager

type LogManager struct {
	msgChan  chan string
	logger   *log.Logger
	file     *os.File
	fileName string
}

func Init(fileName string) bool {
	lm.msgChan = make(chan string, 1000)
	lm.fileName = fileName
	lm.file, lm.logger = newLogger()
	return lm.file == nil || lm.logger == nil
}

//todo 扩展一下write，加上前缀分级
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

func Run() {
	for msg := range lm.msgChan {
		lm.logger.Println(msg)
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
