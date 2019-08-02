package main

import (
	"Emagi/config"
	"Emagi/data"
	"Emagi/log"
	"Emagi/net/tcp"
	"fmt"
	"sync"
	"time"
)

var pbProcessor = &data.PBProcessor{}

func main() {
	//log
	log.Init("./logs/Client")

	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {

		serverConf := config.ServerConf{}
		serverConf.Init("./server_conf.json")
		tcpClient := tcp.NewClient(&serverConf, pbProcessor)
		tcpClient.Start()

		go func(idx int) {
			for j := 0; j < 50; j++ {

				msg := &msg.TestMsg{
					Text: fmt.Sprintf("hello%d", idx),
					Type: 123,
				}
				err := pbProcessor.Serialize(msg)
				if err != nil {
					log.Error(err.Error())
					continue
				}
				tcpClient.WriteMsg(data)

				time.Sleep(2 * time.Second)
			}
		}(i)

		wg.Add(1)
		//test close
		// go func() {
		// 	time.Sleep(8 * time.Second)
		// 	log.Println("close client")
		// 	tcpClient.Close()
		// 	wg.Done()
		// }()
	}
	wg.Wait()
	log.Write("end")
}
