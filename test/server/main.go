package main

import (
	"Emagi/config"
	"Emagi/enet"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	//pprof
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	serverConf := config.ServerConf{}
	serverConf.Init("./server_conf.json")
	tcpServer := enet.TCPServer{}
	tcpServer.Init(&serverConf)
	tcpServer.Start()
}
