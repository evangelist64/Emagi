package main

import (
	"Emagi/enet"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {

		tcpClient := enet.TCPClient{}
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()
			tcpClient.Start()

			for j := 0; j < 50; j++ {
				tcpClient.WriteMsg([]byte(fmt.Sprintf("hello%d", idx)))
				time.Sleep(2 * time.Second)
			}
		}(i)
	}
	wg.Wait()
	log.Println("end")
}
