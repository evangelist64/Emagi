package enet

import (
	"context"
	"log"
	"net"
)

type TCPConn struct {
	conn      net.Conn
	wChan     chan []byte
	closeFlag bool
	belongTo  *TCPServer

	ctx    context.Context
	cancel context.CancelFunc
}

func (p *TCPConn) Destroy() {
	close(p.wChan)
	p.conn.Close()
	p.cancel()

	if p.belongTo != nil {
		p.belongTo.wgConns.Done()
	}
	log.Println("destroy TCPConn")
}

func (p *TCPConn) WriteMsg(b []byte) {
	//已经关闭
	if p.closeFlag {
		return
	}
	//写满了
	if len(p.wChan) == cap(p.wChan) {
		log.Println("wChan full")
		return
	}
	p.wChan <- b
}

func (p *TCPConn) Run() {

	//读循环中止时，不让再写入，写完已有内容后关闭写循环
	defer func() {
		p.wChan <- nil
		p.closeFlag = true
	}()

	//写循环
	go func() {
		defer p.Destroy()
		p.RunWriteLoop()
	}()

	//读循环
	for {
		//todo 通讯协议
		var b [10]byte
		n, err := p.conn.Read(b[:])
		if err != nil {
			log.Println("read error, break")
			break
		}
		if n > 0 {
			log.Println(string(b[:]))
		}

		select {
		case <-p.ctx.Done():
			log.Println("stop read loop")
			return
		default:
		}
	}
}

func (p *TCPConn) RunWriteLoop() {
	for b := range p.wChan {
		//close
		if b == nil {
			log.Println("wChan close sig")
			break
		}

		_, err := p.conn.Write(b)
		if err != nil {
			log.Println("write error, break")
			break
		}
	}
}
