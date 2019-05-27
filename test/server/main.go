package main

import (
	"Emagi/config"
	"Emagi/enet"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	//pprof
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	serverConf := config.ServerConf{}
	serverConf.Init("./server_conf.json")
	tcpServer := enet.TCPServer{}
	tcpServer.Init(&serverConf)

	//test close
	go func() {
		time.Sleep(8 * time.Second)
		log.Println("close server")
		tcpServer.Close()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		tcpServer.Start()
	}()
	wg.Wait()
}
