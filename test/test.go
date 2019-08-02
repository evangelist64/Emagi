package main

import "log"

type TestInterface interface {
	Val() uint32
}

type TestStruct struct {
	Num uint32
}

func (p *TestStruct) Val() uint32 {
	return p.Num
}

func TestFunc(t TestInterface) {
	tt := t.(*TestStruct)
	tt.Num = 2
}

func main() {

	var t = &TestStruct{Num: 1}

	TestFunc(t)

	log.Println(t.Val())

	// ch := make(chan int, 1)
	// ch2 := make(chan int, 1)

	// go func() {
	// 	for i := 0; i < 100; i++ {
	// 		ch <- i
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()

	// time.AfterFunc(10*time.Second, func() {
	// 	ch2 <- 0
	// })

	// for {
	// 	select {
	// 	case i := <-ch:
	// 		log.Printf("ch receive %d", i)
	// 	case i2 := <-ch2:
	// 		log.Printf("ch2 receive %d", i2)
	// 		return
	// 	}
	// }
}
