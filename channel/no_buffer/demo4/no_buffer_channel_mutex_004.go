package main

import (
	"fmt"
	"sync"
)

//传统的基于“共享内存”+“互斥锁”的 Goroutine 安全的计数器的实现

/*
	在这个示例中，我们使用了一个带有互斥锁保护的全局变量作为计数器，
	所有要操作计数器的 Goroutine 共享这个全局变量，
	并在互斥锁的同步下对计数器进行自增操作。
*/

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
