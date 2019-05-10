package main

import (
	"Emagi/enet"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 1; i++ {

		tcpClient := enet.TCPClient{}
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()
			tcpClient.Start()

			for {
				tcpClient.WriteMsg([]byte(fmt.Sprintf("hello%d", idx)))
				time.Sleep(2 * time.Second)
			}
		}(i)
	}
	wg.Wait()
}
