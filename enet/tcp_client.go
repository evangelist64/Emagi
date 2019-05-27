package enet

import (
	"log"
	"net"
	"time"
)

type TCPClient struct {
	conn *TCPConn
}

func (p *TCPClient) Init() {
}

func (p *TCPClient) Start() {

	conn, err := net.DialTimeout("tcp", "127.0.0.1:20000", 5*time.Second)
	if err != nil {
		log.Fatal(err)
		return
	}

	//创建连接
	p.conn = &TCPConn{
		conn:      conn,
		wChan:     make(chan []byte, 100),
		closeFlag: false,
	}
	go p.conn.Run()
}

func (p *TCPClient) WriteMsg(b []byte) {
	p.conn.WriteMsg(b)
}

func (p *TCPClient) Close() {
	p.conn.Destroy()
}
