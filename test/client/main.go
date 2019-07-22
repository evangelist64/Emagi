package main

import (
	"Emagi/config"
	"Emagi/enet"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {

		serverConf := config.ServerConf{}
		serverConf.Init("./server_conf.json")
		tcpClient := enet.NewClient(&serverConf)
		tcpClient.Start()

		go func(idx int) {
			for j := 0; j < 50; j++ {
				tcpClient.WriteMsg([]byte(fmt.Sprintf("hello%d", idx)))
				time.Sleep(2 * time.Second)
			}
		}(i)

		wg.Add(1)
		//test close
		go func() {
			time.Sleep(8 * time.Second)
			log.Println("close client")
			tcpClient.Close()
			wg.Done()
		}()
	}
	wg.Wait()
	log.Println("end")
}
