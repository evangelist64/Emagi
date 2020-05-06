package log

import (
	"Emagi/log"
	"sync"
	"testing"
	"time"
)

func TestLog(t *testing.T) {

	log.Init("Test", true)
	go log.Run()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i < 10000; i++ {
			log.Info("a")
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 10000; i++ {
			log.Info("b")
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 10000; i++ {
			log.Info("c")
		}
		wg.Done()
	}()

	for i := 0; i < 10000; i++ {
		log.Debug("d")
	}
	wg.Wait()

	time.Sleep(10 * time.Second)
}
