package goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func channelDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
		//go worker(i, channels[i])
	}
	for i := 0; i < 10; i++ {
		data := 'a' + i
		channels[i] <- data
	}
	for i := 0; i < 10; i++ {
		data := 'A' + i
		channels[i] <- data
	}
}

func worker(id int, ch chan int) {
	//for {
	//	n, ok := <-ch
	//	if !ok {break}
	//	fmt.Printf("worker %d received %v\n", id, n)
	//}

	// 读完为止
	for n := range ch {
		fmt.Printf("worker %d received %v\n", id, n)
	}
}

/*
chan<- 只收数据的管道
<-chan 只发数据的管道
chan 双向管道
*/
func createWorker(id int) chan<- int {
	ch := make(chan int)
	go func() {
		for {
			fmt.Printf("worker %d received %c\n", id, <-ch)
		}
	}()
	return ch
}

func bufferedChannel() {
	ch := make(chan int, 3)

	go worker(0, ch)

	ch <- 1
	fmt.Println(<-ch)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4 //缓冲区满时会阻塞并painc
}

func channelClose() {
	ch := make(chan int, 3)
	go worker(0, ch)

	ch <- 'a'
	ch <- 'b'
	ch <- 'c'

	close(ch)
}

// deadlock!
func deadlock() {
	//ch := make(chan int, 3)
	//fmt.Println(<- ch)
}

func TestChannel(t *testing.T) {
	channelDemo()

	//bufferedChannel()

	//channelClose()

	for runtime.NumGoroutine() != 1 {
	}
}
