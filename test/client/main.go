package main

import (
	"Emagi/config"
	"Emagi/data"
	"Emagi/log"
	"Emagi/net/tcp"
	"sync"
	"time"
)

var pbProcessor = &data.PBProcessor{}

func main() {

	serverConf := config.ServerConf{}
	serverConf.Init("./server_conf.json")
	//log
	log.Init(serverConf.LogPath+"Client", serverConf.IsDebug)
	go log.Run()
	//pb
	pbProcessor.Init()

	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {

		tcpClient := tcp.NewClient(&serverConf, pbProcessor)
		tcpClient.Start()

		go func(idx int) {
			for j := 0; j < 50; j++ {

				msg := &data.TestMsg{
					Text: "hello",
					Type: int32(idx),
				}
				tcpClient.WriteMsg(msg)

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
