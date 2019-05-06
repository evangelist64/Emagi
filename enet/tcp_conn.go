package enet

import (
	"fmt"
	"net"
)

type TCPConn struct {
	conn net.Conn
}

func (p *TCPConn) Run() {
	fmt.Println("run")
}

func (p *TCPConn) Close() {
	p.conn.Close()
	fmt.Println("close")
}
