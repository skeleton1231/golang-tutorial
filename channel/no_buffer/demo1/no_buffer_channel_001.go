package main

/*
ch1 <- 13    // 将整型字面值13发送到无缓冲channel类型变量ch1中
n := <- ch1  // 从无缓冲channel类型变量ch1中接收一个整型值存储到整型变量n中
ch2 <- 17    // 将整型字面值17发送到带缓冲channel类型变量ch2中
m := <- ch2  // 从带缓冲channel类型变量ch2中接收一个整型值存储到整型变量m中
*/

//channel 是用于 Goroutine 间通信的，所以绝大多数对 channel 的读写都被分别放在了不同的 Goroutine 中。

//无缓冲 channel 类型变量（如 ch1）的发送与接收

func main() {
	/*
		由于无缓冲 channel 的运行时层实现不带有缓冲区，所以 Goroutine 对无缓冲 channel 的接收和发送操作是同步的
		对同一个无缓冲 channel，
		只有对它进行接收操作的 Goroutine 和对它进行发送操作的 Goroutine 都存在的情况下，
		通信才能得以进行，否则单方面的操作会让对应的 Goroutine 陷入挂起状态
	*/
	// ch1 := make(chan int)
	// ch1 <- 13 // fatal error: all goroutines are asleep - deadlock!
	// n := <-ch1
	// println(n)

	ch1 := make(chan int)
	go func() {
		ch1 <- 13 // 将发送操作放入一个新goroutine中执行
	}()
	n := <-ch1
	println(n)

}
