package main

import (
	"Emagi/config"
	"Emagi/log"
	"Emagi/net/msg"
	"Emagi/net/tcp"
	"fmt"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
)

func main() {
	//log
	log.Init("Client")

	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {

		serverConf := config.ServerConf{}
		serverConf.Init("./server_conf.json")
		tcpClient := tcp.NewClient(&serverConf)
		tcpClient.Start()

		go func(idx int) {
			for j := 0; j < 50; j++ {
				msg := &msg.TestMsg{
					Text: fmt.Sprintf("hello%d", idx),
					Type: 123,
				}
				data, _ := proto.Marshal(msg)
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
