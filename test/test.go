package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int, 1)
	ch2 := make(chan int, 1)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
	}()

	time.AfterFunc(10*time.Second, func() {
		ch2 <- 0
	})

	for {
		select {
		case i := <-ch:
			log.Printf("ch receive %d", i)
		case i2 := <-ch2:
			log.Printf("ch2 receive %d", i2)
			return
		}
	}
}
