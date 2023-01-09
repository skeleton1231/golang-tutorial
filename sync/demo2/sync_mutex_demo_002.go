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
	go func(mu *sync.Mutex) {
		mu.Lock()
		i = 10
		time.Sleep(10 * time.Second)
		fmt.Printf("g1: i = %d\n", i)
		mu.Unlock()
		wg.Done()
	}(&mu)

	time.Sleep(time.Second)
	//mu.Lock()

	fmt.Printf("g1: i = %d\n", i)
	//mu.Unlock()

	wg.Wait()
}
