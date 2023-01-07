package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(c chan<- int) {
	var i int = 1
	for {
		time.Sleep(2 * time.Second)
		if len(c) < cap(c) {
			ok := trySend(c, i)
			if ok {
				fmt.Printf("[producer]: send [%d] to channel\n", i)
				i++
				continue
			}
		}
		fmt.Printf("[producer]: try send [%d], but channel is full\n", i)
	}
}

func trySend(c chan<- int, i int) bool {
	select {
	case c <- i:
		return true
	default:
		return false
	}
}

func tryRecv(c <-chan int) (int, bool) {
	select {
	case i := <-c:
		return i, true
	default:
		return 0, false
	}
}

func comsumer(c <-chan int) {
	for {
		if len(c) > 0 {
			i, ok := tryRecv(c)
			if !ok {
				fmt.Printf("[producer]: try send [%d], but channel is full\n", i)
				time.Sleep(1 * time.Second)
				continue
			}
			fmt.Printf("[consumer]: recv [%d] from channel\n", i)
			if i >= 3 {
				fmt.Println("[consumer]: exit")
				return
			}
		}

	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int, 3)
	wg.Add(2)
	go func() {
		producer(c)
		wg.Done()
	}()

	go func() {
		comsumer(c)
		wg.Done()
	}()
	wg.Wait()
}
