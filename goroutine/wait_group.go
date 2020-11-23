package main

import (
	"fmt"
	"sync"
)

func workerWait(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("worker %d: %d\n", x, i)
	}
}

func main() {
	//done := make(chan bool)
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		fmt.Println(i)
	//	}
	//	done<- true
	//}()
	//<-done

	var wg sync.WaitGroup

	wg.Add(2)
	go workerWait(1, &wg)
	go workerWait(2, &wg)

	wg.Wait()
}
