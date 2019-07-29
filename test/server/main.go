package main

import (
	"Emagi/config"
	"Emagi/log"
	"Emagi/net/tcp"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {
	//log
	log.Init("Server")

	wg := sync.WaitGroup{}
	//pprof host
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	serverConf := config.ServerConf{}
	serverConf.Init("./server_conf.json")
	tcpServer := tcp.NewServer(&serverConf)

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
