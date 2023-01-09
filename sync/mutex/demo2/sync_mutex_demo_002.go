package main

import (
	"fmt"
	"sync"
	"time"
)

// 我们推荐通过闭包方式，或者是传递类型实例（或包裹该类型的类型实例）的地址（指针）的方式进行
func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	i := 0
	wg.Add(1)

	/*
		一旦某个 Goroutine 调用的 Mutex 执行 Lock 操作成功，它将成功持有这把互斥锁。
		这个时候，如果有其他 Goroutine 执行 Lock 操作，
		就会阻塞在这把互斥锁上，直到持有这把锁的 Goroutine 调用 Unlock 释放掉这把锁后，
		才会抢到这把锁的持有权并进入临界区
	*/
	go func(mu *sync.Mutex) {
		mu.Lock()
		i = 10
		time.Sleep(10 * time.Second)
		fmt.Printf("g1: i = %d\n", i)
		mu.Unlock()
		wg.Done()
	}(&mu)

	time.Sleep(time.Second)
	mu.Lock()
	i = 1
	fmt.Printf("g0: i = %d\n", i)
	mu.Unlock()
	wg.Wait()

}
