package service

//服务抽象
type Service struct {
	ch chan func() //主逻辑channel，业务逻辑全都同步执行
}

func (p *Service) Init() {
	p.ch = make(chan func(), 10)
}

func (p *Service) Run() {
	for {
		f := <-p.ch
		f()
	}
}

func (p *Service) (f func()) {
	p.ch <- f
}
