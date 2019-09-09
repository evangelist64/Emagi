package main

import (
	"Emagi/config"
	"Emagi/data"
	"Emagi/log"
	"Emagi/net/tcp"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"sync"

	"github.com/golang/protobuf/proto"
)

var pbProcessor = &data.PBProcessor{}

func OnTestMsg(p proto.Message) {
	testMsg := p.(*data.TestMsg)
	log.Info(testMsg.GetText() + strconv.Itoa((int)(testMsg.GetType())))
}

func main() {
	//pprof host
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	serverConf := config.ServerConf{}
	serverConf.Init("./server_conf.json")

	//log
	log.Init(serverConf.LogPath+"/Server", serverConf.IsDebug)
	go log.Run()

	//processor init
	pbProcessor.Init()
	pbProcessor.Register((*data.TestMsg)(nil), OnTestMsg)

	wg := sync.WaitGroup{}

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
