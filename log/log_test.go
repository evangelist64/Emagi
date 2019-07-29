package log

import (
	"Emagi/log"
	"sync"
	"testing"
)

func TestLog(t *testing.T) {

	log.Init("Test")

	for i := 0; i < 100000; i++ {
		log.Debug("ffffffffffffffggggggggggggg")
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i < 100000; i++ {
			log.Info("ffffffffffffffggggggggggggg")
		}
		wg.Done()
	}()
	wg.Wait()
}
