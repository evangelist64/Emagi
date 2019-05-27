package enet

import (
	"Emagi/config"
	"context"
	"log"
	"net"
	"sync"
	"time"
)

type TCPServer struct {
	conf     *config.ServerConf
	listener net.Listener

	ctx    context.Context
	cancel context.CancelFunc

	wgConns *sync.WaitGroup
	// wgAccept *sync.WaitGroup
}

func (p *TCPServer) Init(conf *config.ServerConf) {

	p.conf = conf
	ctx, cancel := context.WithCancel(context.Background())
	p.ctx = ctx
	p.cancel = cancel
	p.wgConns = &sync.WaitGroup{}
	// p.wgAccept = &sync.WaitGroup{}
}

func (p *TCPServer) Start() {
	listener, err := net.Listen("tcp", p.conf.Address)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listen on %s", p.conf.Address)
	p.listener = listener

	// p.wgAccept.Add(1)

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
			log.Printf("accept error, tcpServer exit")
			return
		}
		tempDelay = 0

		p.wgConns.Add(1)
		//创建连接
		tcpConn := &TCPConn{
			conn:      conn,
			wChan:     make(chan []byte, 100),
			closeFlag: false,
			belongTo:  p,
		}
		tcpConn.ctx, tcpConn.cancel = context.WithCancel(p.ctx)

		go tcpConn.Run()

		select {
		case <-p.ctx.Done():
			log.Println("stop accept loop")
			return
		default:
		}
	}
}

func (p *TCPServer) Close() {
	p.cancel()
	p.listener.Close()

	p.wgConns.Wait()
	log.Printf("all conn closed")
}
