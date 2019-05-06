package enet

import (
	"Emagi/config"
	"log"
	"net"
	"sync"
	"time"
)

type TCPServer struct {
	conf     *config.ServerConf
	listener net.Listener

	conns      map[TCPConn]struct{}
	connsMutex sync.Mutex
}

func (p *TCPServer) Init(conf *config.ServerConf) {
	p.conf = conf
	p.conns = make(map[TCPConn]struct{})
}

func (p *TCPServer) Start() {
	listener, err := net.Listen("tcp", p.conf.Address)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listen on %s", p.conf.Address)
	p.listener = listener

	//接受客户端连接
	var tempDelay time.Duration
	for {
		conn, err := p.listener.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				time.Sleep(tempDelay)
				continue
			}
			//todo log
			return
		}
		tempDelay = 0

		tcpConn := new(TCPConn)
		tcpConn.conn = conn

		p.connsMutex.Lock()
		if len(p.conns) >= 10000 {
			p.connsMutex.Unlock()
			conn.Close()
			continue
		}
		p.conns[*tcpConn] = struct{}{}
		p.connsMutex.Unlock()

		go func() {

			defer func() {
				p.connsMutex.Lock()
				delete(p.conns, *tcpConn)
				p.connsMutex.Unlock()

				tcpConn.Close()
			}()
			tcpConn.Run()
		}()
	}
}

func (p *TCPServer) Close() {
	p.connsMutex.Lock()
	for tcpConn := range p.conns {
		tcpConn.Close()
	}
	p.conns = nil
	p.connsMutex.Unlock()
}
