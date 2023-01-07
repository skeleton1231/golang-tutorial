package demo2_test

import "testing"

var c1 chan string
var c2 chan string

func init() {
	c1 = make(chan string, 100)
	for i := 0; i < 1; i++ {
		go func() {
			for {
				<-c1 // recive
			}
		}()
	}
	c2 = make(chan string, 100)
	for i := 0; i < 1; i++ {
		go func() {
			for {
				c2 <- "hello"
			}
		}()
	}
}

func send(msg string) {
	c1 <- msg
}

func recv() {
	<-c2
}

func BenchmarkBufferedChan1To1SendCap100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		send("hello")
	}
}
func BenchmarkBufferedChan1To1RecvCap100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		recv()
	}
}
