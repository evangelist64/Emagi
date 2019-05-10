package enet

import (
	"fmt"
	"net"
)

type TCPConn struct {
	conn  net.Conn
	wChan chan []byte
}

func (p *TCPConn) Init(conn net.Conn) {
	p.conn = conn
	p.wChan = make(chan []byte, 100)
}

func (p *TCPConn) Run() {
	fmt.Println("run")
}

func (p *TCPConn) Close() {
	p.conn.Close()
	fmt.Println("close")
}

func (p *TCPConn) WriteMsg(b []byte) {
	p.wChan <- b
}

func (p *TCPConn) RunReadLoop() {
	for {
		var b [10]byte
		n, err := p.conn.Read(b[:])
		if err != nil {
			break
		}
		if n > 0 {
			fmt.Println(string(b[:]))
		}
	}
}

func (p *TCPConn) RunWriteLoop() {

	for b := range p.wChan {
		if b == nil {
			break
		}

		_, err := p.conn.Write(b)
		if err != nil {
			break
		}
	}
}
