package main

import (
	"fmt"
	"math/rand"
	"time"
)

//
func genChain() <-chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

//
func createWorker2(id int) chan<- int {
	ch := make(chan int)
	go func() {
		for {
			fmt.Printf("worker %d received %v\n", id, <-ch)
		}
	}()
	return ch
}

func main() {
	over := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	var p1, p2 <-chan int = genChain(), genChain() //生产消息
	consumer := createWorker2(0)                   // 消费消息
	n := 0
	hasValue := false
	// 多路复用
	for {
		var activeConsumer chan<- int
		if hasValue {
			activeConsumer = consumer
		}
		select {
		case n = <-p1:
			hasValue = true
		case n = <-p2:
			hasValue = true
		case activeConsumer <- n:
			hasValue = false
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("111")
		case <-over:
			fmt.Println("bye")
			return //结束函数
			//default:
			//	fmt.Println("No value received")
		}
	}
}
