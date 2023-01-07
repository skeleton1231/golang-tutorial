package main

import (
	"fmt"
	"time"
)

// 第一种用法：用作信号传递
// 1 对 1 通知信号这种情况
/*
	在这个例子中，spawn 函数返回的 channel，被用于承载新 Goroutine 退出的“通知信号”，
	这个信号专门用作通知 main goroutine。
	main goroutine 在调用 spawn 函数后一直阻塞在对这个“通知信号”的接收动作上
*/

type signal struct{}

func worker() {
	fmt.Println("worker is working...")
	time.Sleep(1 * time.Second)
}

func spawn(f func()) chan signal {
	c := make(chan signal)
	go func() {
		fmt.Println("worker start to work...")
		f()
		c <- signal{}
	}()
	return c
}

func main() {
	fmt.Println("start a worker...")
	c := spawn(worker)
	<-c
	fmt.Println("worker work done!")
}
