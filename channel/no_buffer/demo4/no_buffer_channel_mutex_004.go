package main

import (
	"fmt"
	"sync"
)

//传统的基于“共享内存”+“互斥锁”的 Goroutine 安全的计数器的实现

type counter struct {
	sync.Mutex
	i int
}

var cter counter

func Increase() int {
	cter.Lock()
	defer cter.Unlock()
	cter.i++
	return cter.i
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := Increase()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
