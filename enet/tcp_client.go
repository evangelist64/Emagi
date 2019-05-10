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

	tcpConn := new(TCPConn)
	tcpConn.Init(conn)
	p.conn = tcpConn

	go func() {
		defer tcpConn.Close()
		tcpConn.RunReadLoop()
	}()

	go func() {
		defer tcpConn.Close()
		tcpConn.RunWriteLoop()
	}()
}

func (p *TCPClient) WriteMsg(b []byte) {
	p.conn.WriteMsg(b)
}

func (p *TCPClient) Close() {
	p.conn.Close()
}
