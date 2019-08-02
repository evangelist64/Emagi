package main

import (
	"Emagi/config"
	"Emagi/data"
	"Emagi/log"
	"Emagi/net/tcp"
	"net/http"
	_ "net/http/pprof"
	"sync"

	"github.com/golang/protobuf/proto"
)

var pbProcessor = &data.PBProcessor{}

func OnTestMsg(p proto.Message) {
	testMsg := p.(*msg.TestMsg)
	log.Info(testMsg.GetText())
}

func main() {
	//pprof host
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	//log
	log.Init("./logs/Server")

	//processor init
	pbProcessor.Register((*msg.TestMsg)(nil), OnTestMsg)

	wg := sync.WaitGroup{}

	serverConf := config.ServerConf{}
	serverConf.Init("./server_conf.json")
	tcpServer := tcp.NewServer(&serverConf, pbProcessor)

	//test close
	// go func() {
	// 	time.Sleep(8 * time.Second)
	// 	log.Println("close server")
	// 	tcpServer.Close()
	// }()

	wg.Add(1)
	go func() {
		defer wg.Done()
		tcpServer.Start()
	}()
	wg.Wait()
}
