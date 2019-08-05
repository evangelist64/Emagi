package tcp

import (
	"Emagi/config"
	"Emagi/data"
	"Emagi/log"
	"context"
	"net"
	"sync"
	"time"
)

type TCPClient struct {
	conn *TCPConn           //todo改成支持多个连接
	conf *config.ServerConf //todo改成客户端配置

	dp data.DataProcessor

	ctx    context.Context
	cancel context.CancelFunc
}

func NewClient(conf *config.ServerConf, dp data.DataProcessor) *TCPClient {
	c := &TCPClient{
		conf: conf,
		conn: nil,
	}
	c.dp = dp
	c.ctx, c.cancel = context.WithCancel(context.Background())
	return c
}

func (p *TCPClient) Start() {

	//TODO 尝试重连
	conn, err := net.DialTimeout("tcp", p.conf.Address, 5*time.Second)
	if err != nil {
		log.Error(err.Error())
		return
	}

	//创建连接
	tcpConn := &TCPConn{
		Id:       0,
		conn:     conn,
		wChan:    make(chan interface{}, 100),
		wg:       &sync.WaitGroup{},
		wgParent: nil,
		dp:       p.dp,
	}
	tcpConn.ctx, tcpConn.cancel = context.WithCancel(p.ctx)
	p.conn = tcpConn

	go p.conn.Run()
}

func (p *TCPClient) WriteMsg(msg interface{}) {
	p.conn.WriteMsg(msg)
}

func (p *TCPClient) Close() {
	p.cancel()
}
