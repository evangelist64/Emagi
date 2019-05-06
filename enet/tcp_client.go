package enet

import (
	"net"
	"log"
	"time"
)

type TCPClient struct {
	conn net.Conn
}

func (p *TCPClient)Init() {
}

func (p *TCPClient)Start(){
	for {
		conn, err := net.DialTimeout("tcp", "127.0.0.1:20000",5*time.Second)
		if err != nil {
			log.Fatal(err)
			continue
		}

		p.conn = conn
		break
	}
	tcpConn := new(TCPConn)
	tcpConn.conn = p.conn
	tcpConn.Run()
}

func (p *TCPClient)Close(){
	p.conn.Close()
}