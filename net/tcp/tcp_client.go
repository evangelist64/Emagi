package tcp

import (
	"Emagi/config"
	"Emagi/log"
	"context"
	"net"
	"sync"
	"time"
)

type TCPClient struct {
	conn *TCPConn           //todo改成支持多个连接
	conf *config.ServerConf //todo改成客户端配置

	ctx    context.Context
	cancel context.CancelFunc
}

func NewClient(conf *config.ServerConf) *TCPClient {
	c := &TCPClient{
		conf: conf,
		conn: nil,
	}
	c.ctx, c.cancel = context.WithCancel(context.Background())
	return c
}

func (p *TCPClient) Start() {

	//TODO 尝试重连
	conn, err := net.DialTimeout("tcp", p.conf.Address, 5*time.Second)
	if err != nil {
		log.Write(err.Error())
		return
	}

	//创建连接
	tcpConn := &TCPConn{
		Id:       0,
		conn:     &conn,
		wChan:    make(chan []byte, 100),
		wg:       &sync.WaitGroup{},
		wgParent: nil,
	}
	tcpConn.ctx, tcpConn.cancel = context.WithCancel(p.ctx)
	p.conn = tcpConn

	go p.conn.Run()
}

func (p *TCPClient) WriteMsg(b []byte) {
	p.conn.WriteMsg(b)
}

func (p *TCPClient) Close() {
	p.cancel()
}
