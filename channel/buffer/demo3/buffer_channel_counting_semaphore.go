package main

import (
	"log"
	"sync"
	"time"
)

/*
	Go 并发设计的一个惯用法，就是将带缓冲 channel 用作计数信号量（counting semaphore）。
	带缓冲 channel 中的当前数据个数代表的是，
	当前同时处于活动状态（处理业务）的 Goroutine 的数量，
	而带缓冲 channel 的容量（capacity），
	就代表了允许同时处于活动状态的 Goroutine 的最大数量。
	向带缓冲 channel 的一个发送操作表示获取一个信号量，
	而从 channel 的一个接收操作则表示释放一个信号量。
*/

/*
	这个示例使用了一个容量（capacity）为 3 的带缓冲 channel: active 作为计数信号量，
	这意味着允许同时处于活动状态的最大 Goroutine 数量为 3
*/

var active = make(chan struct{}, 3)
var jobs = make(chan int, 10)

func main() {
	go func() {
		for i := 0; i < 8; i++ {
			jobs <- (i + 1)
		}
		close(jobs) //发送端close
	}()

	var wg sync.WaitGroup
	for j := range jobs {
		wg.Add(1)
		go func(j int) {
			active <- struct{}{}
			log.Printf("handle job: %d\n", j)
			time.Sleep(2 * time.Second)
			<-active
			wg.Done()
		}(j)
	}
	wg.Wait()
}
