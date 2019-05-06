package main

import (
	"Emagi/config"
	"Emagi/enet"
)

func main() {
	serverConf := config.ServerConf{}
	serverConf.Init("./server_conf.json")
	tcpServer := enet.TCPServer{}
	tcpServer.Init(&serverConf)

	tcpServer.Start()
}
