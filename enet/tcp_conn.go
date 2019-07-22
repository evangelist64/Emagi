package enet

import (
	"context"
	"log"
	"net"
	"sync"
)

type TCPConn struct {
	Id       uint32
	conn     *net.Conn
	wChan    chan []byte //写channel
	wg       *sync.WaitGroup
	wgParent *sync.WaitGroup

	ctx    context.Context
	cancel context.CancelFunc
}

func (p *TCPConn) Destroy() {
	close(p.wChan)
	(*p.conn).Close()
	log.Println("destroy TCPConn")
}

func (p *TCPConn) WriteMsg(b []byte) {
	select {
	//已经关闭,不让再写进去
	case <-p.ctx.Done():
		return
	default:
		//写满了
		if len(p.wChan) == cap(p.wChan) {
			log.Println("wChan full, send failed")
			return
		}
		p.wChan <- b
	}
}

func (p *TCPConn) Run() {
	//写循环
	p.wg.Add(1)
	go func() {
		ctxw, cancelw := context.WithCancel(p.ctx)

		defer func() {
			cancelw()
			//出问题退出写循环，要通知上层goroutine执行退出操作
			p.cancel()
			p.wg.Done()
		}()

		for {
			select {
			case <-ctxw.Done():
				return
			default:
				b := <-p.wChan
				//close
				if b == nil {
					log.Println("wChan close sig")
					return
				}

				_, err := (*p.conn).Write(b)
				if err != nil {
					log.Println("write error, break")
					return
				}
			}
		}
	}()

	defer func() {
		//如果是读操作出问题退出读循环，需要调用cancel让子协程退出
		p.cancel()

		//等待子协程退出
		p.wg.Wait()
		p.Destroy()
		if p.wgParent != nil {
			p.wgParent.Done()
		}
	}()
	//读循环
	for {
		select {
		case <-p.ctx.Done():
			log.Println("ctx done, read loop stop")
			return
		default:
			var b [10]byte
			n, err := (*p.conn).Read(b[:])
			if err != nil {
				log.Println("read error, read loop stop")
				return
			}
			if n > 0 {
				log.Println(string(b[:]))
			}
		}
	}
}
