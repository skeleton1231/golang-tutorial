package main

import (
	"fmt"
	"sync"
	"time"
)

func produce(ch chan<- int) { // send only
	for i := 0; i < 10; i++ {
		ch <- i + 1
		time.Sleep(time.Second)
	}
	close(ch) //important!!! channel 的一个使用惯例，那就是发送端负责关闭 channel
}

func comsume(ch <-chan int) { // read only
	for n := range ch {
		fmt.Println(n)
	}
}

func main() {
	/*
		和无缓冲 channel 相反，带缓冲 channel 的运行时层实现带有缓冲区，
		因此，对带缓冲 channel 的发送操作在缓冲区未满、
		接收操作在缓冲区非空的情况下是异步的（发送或接收不需要阻塞等待）。


		ch2 := make(chan int, 1)
		n := <-ch2 // 由于此时ch2的缓冲区中无数据，因此对其进行接收操作将导致goroutine挂起

		ch3 := make(chan int, 1)
		ch3 <- 17  // 向ch3发送一个整型数17
		ch3 <- 27  // 由于此时ch3中缓冲区已满，再向ch3发送数据也将导致goroutine挂起
	*/
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		produce(ch)
		wg.Done()
	}()

	go func() {
		comsume(ch)
		wg.Done()
	}()

	wg.Wait()
}
