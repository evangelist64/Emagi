package tcp

import (
	"Emagi/config"
	"Emagi/data"
	"Emagi/log"
	"context"
	"fmt"
	"net"
	"sync"
	"time"
)

type TCPServer struct {
	conf     *config.ServerConf
	listener net.Listener

	conns   *sync.Map       //add和遍历操作不在同个协程上
	wgConns *sync.WaitGroup //等待子协程结束

	dp data.DataProcessor

	ctx    context.Context
	cancel context.CancelFunc

	curConnId uint32
}

func NewServer(conf *config.ServerConf, dp data.DataProcessor) *TCPServer {
	s := &TCPServer{
		conf:      conf,
		conns:     &sync.Map{},
		wgConns:   &sync.WaitGroup{},
		curConnId: 0,
	}
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.dp = dp
	return s
}

func (p *TCPServer) getIncConnId() uint32 {
	p.curConnId++
	return p.curConnId
}

func (p *TCPServer) Run() {
	listener, err := net.Listen("tcp", p.conf.Address)
	if err != nil {
		log.Write(err.Error())
	}
	log.Info(fmt.Sprintf("listen on %s", p.conf.Address))
	p.listener = listener

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
			log.Error("accept error, tcpServer exit")
			return
		}
		tempDelay = 0.

		p.wgConns.Add(1)
		//创建连接
		tcpConn := &TCPConn{
			Id:       p.getIncConnId(),
			conn:     conn,
			wChan:    make(chan interface{}, 100),
			wg:       &sync.WaitGroup{},
			wgParent: p.wgConns,
			dp:       p.dp,
		}
		tcpConn.ctx, tcpConn.cancel = context.WithCancel(p.ctx)
		p.conns.Store(tcpConn.Id, tcpConn)

		go tcpConn.Run()
	}
}

func (p *TCPServer) Close() {
	p.listener.Close()
	p.cancel()

	p.wgConns.Wait()
	log.Info("all conns closed")
}
