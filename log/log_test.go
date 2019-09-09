package log

import (
	"Emagi/log"
	"testing"
)

func TestLog(t *testing.T) {

	log.Init("Test", true)
	go log.Run()

	go func() {
		for i := 0; i < 100000; i++ {
			log.Info("b")
		}
	}()

	go func() {
		for i := 0; i < 100000; i++ {
			log.Info("c")
		}
	}()

	for i := 0; i < 100000; i++ {
		log.Debug("a")
	}

}
